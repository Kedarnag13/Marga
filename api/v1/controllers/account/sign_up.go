package account

import (
	"encoding/json"
	// "github.com/Kedarnag13/Marga/api/v1/config/db"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

type registrationController struct{}

var Registration registrationController

func (reg registrationController) Create(rw http.ResponseWriter, req *http.Request) {
	db, err := gorm.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	defer db.Close()

	var user models.User
	var session models.Session
	var device models.Device

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	check_mobile_number_exists := db.First(&user, user.MobileNumber)
	if check_mobile_number_exists.RecordNotFound() != true {
		b, err := json.Marshal(models.Response{
			Message: "",
			Error:   "MobileNumber already exists!",
			Success: false,
		})
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		if user.Password != user.PasswordConfirmation {
			b, err := json.Marshal(models.Response{
				Message: "",
				Error:   "Password and PasswordConfirmation do not match!",
				Success: false,
			})
			if err != nil {
				panic(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto end
		} else {
			check_session_exists_for_mobile_number := db.First(&session, session.DeviseToken)
			if check_session_exists_for_mobile_number.RecordNotFound() == true {
				key := []byte("traveling is fun")
				password := []byte(user.Password)
				confirm_password := []byte(user.PasswordConfirmation)
				encrypt_password := controllers.Encrypt(key, password)
				encrypt_password_confirmation := controllers.Encrypt(key, confirm_password)

				user = models.User{Name: user.Name, Email: user.Email, MobileNumber: user.MobileNumber, Latitude: user.Latitude, Longitude: user.Longitude, Password: encrypt_password, PasswordConfirmation: encrypt_password_confirmation, City: user.City, DeviseToken: user.DeviseToken, Type: user.Type}
				db.Create(&user)
				session = models.Session{UserId: user.ID, DeviseToken: user.DeviseToken}
				db.Create(&session)
				device = models.Device{UserId: user.ID, SessionId: session.ID}
				db.Create(&device)

				b, err := json.Marshal(models.Response{
					Message: "User and Session created Successfully!",
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
	}
end:
}
