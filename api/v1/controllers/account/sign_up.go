package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/asaskevich/govalidator"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type registrationController struct{}

var Registration registrationController

func (r registrationController) Create(rw http.ResponseWriter, req *http.Request) {

	var u models.User

	flag := 1

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}

	check_mobile_number, err := db.Query("SELECT mobile_number from users")
	if err != nil {
		log.Fatal(err)
	}

	fetch_id, err := db.Query("SELECT coalesce(max(id), 0) FROM users")
	if err != nil {
		log.Fatal(err)
	}
	if flag == 1 {
		email := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		exp, err := regexp.Compile(email)
		if err != nil {
			os.Exit(1)
		}
		if u.Name == "" || u.Username == "" || u.Email == "" || !exp.MatchString(u.Email) || u.Mobile_number == "" || u.Password == "" || u.Password_confirmation == "" || u.Devise_token == "" || u.City == "" || u.Type == "" {

			result, err := govalidator.ValidateStruct(u)
			if err != nil {
				println("error: " + err.Error())
			}
			fmt.Println(result)
			flag = 0
			b, err := json.Marshal(models.ErrorMessage{
				Success: "false",
				Error:   err.Error(),
			})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto create_user_end
		}
	}

	if flag == 1 {
		for check_mobile_number.Next() {
			var mobile_number string
			err = check_mobile_number.Scan(&mobile_number)
			if err != nil {
				log.Fatal(err)
			}

			if mobile_number == u.Mobile_number {
				b, err := json.Marshal(models.ErrorMessage{
					Success: "false",
					Error:   "Mobile number already exist",
				})
				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				fmt.Println("Mobile number already exist")
				flag = 0
				goto create_user_end
			}
		}
	}
	defer check_mobile_number.Close()

	if u.Password != u.Password_confirmation {
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "Password and Password_confirmation do not match",
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto create_user_end
	}

	if flag == 1 {
		session_response, err := db.Query("SELECT devise_token,user_id from sessions")
		if err != nil {
			log.Fatal(err)
		}
		for session_response.Next() {
			var devise_token string
			var id int
			err := session_response.Scan(&devise_token, &id)
			if err != nil {
				log.Fatal(err)
			}
			if devise_token == u.Devise_token {
				b, err := json.Marshal(models.ErrorMessage{
					Success: "false",
					Error:   "Session already Exist",
				})

				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				fmt.Println("Session already Exist")
				goto create_user_end
			}
		}
	}

	if flag == 1 {
		for fetch_id.Next() {
			var id int
			err = fetch_id.Scan(&id)

			if err != nil {
				log.Fatal(err)
			}
			id = id + 1

			var insert_user string = "insert into users (id, name, username, email, mobile_number, latitude, longitude, password, password_confirmation, city, device_token, ward_id, type) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
			db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
			if err != nil {
				log.Fatal(err)
			}
			prepare_insert_user, err := db.Prepare(insert_user)
			if err != nil {
				log.Fatal(err)
			}

			key := []byte("traveling is fun")
			password := []byte(u.Password)
			confirm_password := []byte(u.Password_confirmation)
			encrypt_password := controllers.Encrypt(key, password)
			encrypt_password_confirmation := controllers.Encrypt(key, confirm_password)

			user_res, err := prepare_insert_user.Exec(id, u.Name, u.Username, u.Email, u.Mobile_number, u.Latitude, u.Longitude, encrypt_password, encrypt_password_confirmation, u.City, u.Devise_token, u.Ward_id, u.Type)
			if err != nil || user_res == nil {
				log.Fatal(err)
			}

			var devise string = "insert into devices(devise_token,user_id)values ($1,$2)"
			dev, err := db.Prepare(devise)
			if err != nil {
				log.Fatal(err)
			}
			dev_res, err := dev.Exec(u.Devise_token, id)
			if err != nil || dev_res == nil {
				log.Fatal(err)
			}
			var session string = "insert into sessions (user_id,devise_token) values ($1,$2)"
			ses, err := db.Prepare(session)
			if err != nil {
				log.Fatal(err)
			}
			session_res, err := ses.Exec(id, u.Devise_token)
			if err != nil || session_res == nil {
				log.Fatal(err)
			}
			fmt.Println("User created Successfully!")

			user := models.User{id, u.Name, u.Username, u.Email, u.Mobile_number, u.Latitude, u.Longitude, u.Password, u.Password_confirmation, u.City, u.Devise_token, u.Ward_id, u.Type}

			b, err := json.Marshal(models.SuccessfulSignIn{
				Success: "true",
				Message: "User created Successfully!",
				User:    user,
				Session: models.SessionDetails{id, u.Devise_token},
			})

			if err != nil {
				log.Fatal(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			defer prepare_insert_user.Close()
		}
		defer fetch_id.Close()
	}
create_user_end:
	db.Close()
}
