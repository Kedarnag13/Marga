package account

import (
	"encoding/json"
	// "github.com/Kedarnag13/Marga/api/v1/config/db"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	defer db.Close()

	var user models.User
	var session models.Session
	var device models.Device

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &session)
	if err != nil {
		panic(err)
	}

	var id uint
	var password string
	find_user := db.First(&user, "mobile_number = ?", session.User.MobileNumber)
	if find_user.RecordNotFound() == true {
		b, err := json.Marshal(models.Response{
			Message: "",
			Error:   "User does not exist!",
			Success: false,
		})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		find_user := db.Table("users").Where("mobile_number = ?", session.User.MobileNumber).Select("id", "password").Row()
		find_user.Scan(&id, &password)
		check_session_exists := db.Where("devise_token = ?", session.DeviseToken).First(&session)
		if check_session_exists.RecordNotFound() == true {
			key := []byte("traveling is fun")
			log.Println(password)
			db_password := password
			decrypt_password := controllers.Decrypt(key, db_password)
			log.Println(decrypt_password, user.Password)
			if decrypt_password == user.Password {
				session = models.Session{UserId: id, DeviseToken: session.DeviseToken}
				db.Create(&session)
				device = models.Device{DeviseToken: session.DeviseToken, UserId: id, SessionId: session.ID}
				db.Create(&device)

				b, err := json.Marshal(models.Response{
					Message: "Session created Successfully!",
					Error:   "",
					Success: true,
				})

				if err != nil {
					panic(err)
				}

				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				goto end
			} else {
				b, err := json.Marshal(models.Response{
					Message: "",
					Error:   "Invalid MobileNumber or Password!",
					Success: false,
				})

				if err != nil {
					panic(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				goto end
			}
		} else {
			b, err := json.Marshal(models.Response{
				Message: "",
				Error:   "Session already exists!",
				Success: false,
			})

			if err != nil {
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto end
		}
	}
end:
}

func (s *sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var session models.Session
	vars := mux.Vars(req)
	devisetoken := vars["devisetoken"]
	session.DeviseToken = devisetoken

	user := db.First(&session, session.DeviseToken)
	if user.RecordNotFound() == true {
		b, err := json.Marshal(models.Response{
			Message: "",
			Error:   "User does not exist!",
			Success: false,
		})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		db.Exec("DELETE FROM sessions WHERE devise_token = ?", session.DeviseToken)
		db.Exec("DELETE FROM devices WHERE devise_token = ?", session.DeviseToken)
		b, err := json.Marshal(models.Response{
			Message: "Session destroyed Successfully!",
			Error:   "",
			Success: true,
		})

		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	}
end:
}
