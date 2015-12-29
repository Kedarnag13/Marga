package account

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	var u models.User

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}

	response, error, user := session(u, true, false)

	if error == true {
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   response,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		b, err := json.Marshal(models.SuccessfulSignIn{
			Success: "true",
			Message: "Logged in Successfully",
			User:    user,
			Session: models.SessionDetails{user.Id, user.Devise_token},
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
end:
}

func (s *sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {

	var u models.User
	vars := mux.Vars(req)
	devise_token := vars["devise_token"]
	u.Devise_token = devise_token

	response, error, user := session(u, false, true)

	if error == true {
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   response,
		})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		b, err := json.Marshal(models.Message{
			User:    user,
			Success: "true",
			Message: response,
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
end:
}

func session(user models.User, login, logout bool) (string, bool, models.User) {

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	get_session, err := db.Query("SELECT * from sessions where devise_token=$1", user.Devise_token)
	if err != nil {
		log.Fatal(err)
	}
	get_user_id, err := db.Query("SELECT id FROM users WHERE mobile_number=$1", user.Mobile_number)
	if err != nil {
		log.Fatal(err)
	}
	flag := 0
	for get_user_id.Next() {
		var id int
		get_user_id.Scan(&id)
		user_existance := controllers.Check_for_user(id)
		if user_existance == false {
			return "User Does not exist", true, user
		}
	}
	defer get_user_id.Close()
	if login == true {
		if user.Mobile_number == "" || user.Password == "" {
			return "Please enter credentials to log in", true, user
		}
	}
	if user.Devise_token == "" {
		return "Devise Token can't be empty", true, user
	} else {
		for get_session.Next() {
			flag = 1
			delete_sessions, err := db.Prepare("DELETE from sessions where devise_token =$1")
			delete_sessions_res, err := delete_sessions.Exec(user.Devise_token)
			if err != nil || delete_sessions_res == nil {
				log.Fatal(err)
			}

			delete_devise, err := db.Prepare("DELETE from devices where devise_token =$1")
			delete_devise_res, err := delete_devise.Exec(user.Devise_token)
			if err != nil || delete_devise_res == nil {
				log.Fatal(err)
			}
			if logout == true {
				user := models.User{0, "", "", "", "", 0, 0, "", "", "", user.Devise_token, 0, ""}
				return "Logged out Successfully", false, user
			}
		}
		defer get_session.Close()
		if logout == true && flag == 0 {
			return "Session does not exist", true, user
		}
		if login == true {
			get_user, err := db.Query("SELECT id,name,username, email, mobile_number, latitude, longitude, password, password_confirmation, city, device_token, ward_id, type FROM users WHERE mobile_number=$1", user.Mobile_number)
			if err != nil {
				log.Fatal(err)
			}
			for get_user.Next() {
				var id int
				var name string
				var username string
				var email string
				var mobile_number string
				var latitude float64
				var longitude float64
				var password string
				var password_confirmation string
				var city string
				var ward_id int
				var devise_token string
				var user_type string

				err := get_user.Scan(&id, &name, &username, &email, &mobile_number, &latitude, &longitude, &password, &password_confirmation, &city, &devise_token, &ward_id, &user_type)
				if err != nil {
					log.Fatal(err)
				}
				key := []byte("traveling is fun")
				db_password := password
				decrypt_password := controllers.Decrypt(key, db_password)
				if mobile_number == user.Mobile_number && decrypt_password == user.Password {
					var devise string = "insert into devices(devise_token,user_id)values ($1,$2)"
					dev, err := db.Prepare(devise)
					if err != nil {
						log.Fatal(err)
					}
					dev_res, err := dev.Exec(user.Devise_token, id)
					if err != nil || dev_res == nil {
						log.Fatal(err)
					}
					defer dev.Close()

					var session string = "insert into sessions (user_id,devise_token) values ($1,$2)"
					ses, err := db.Prepare(session)
					if err != nil {
						log.Fatal(err)
					}
					res, err := ses.Exec(id, user.Devise_token)
					if err != nil || res == nil {
						log.Fatal(err)
					}
					user_details := models.User{id, name, username, email, mobile_number, latitude, longitude, "", "", city, devise_token, ward_id, user_type}
					return "Logged in Successfully", false, user_details
				}
			}
			defer get_user.Close()
		}
	}
	db.Close()
	return "Invalid Mobile Number or Password", true, user
}
