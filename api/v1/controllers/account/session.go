package account

import (
	"encoding/json"
	// "github.com/Kedarnag13/Marga/api/v1/config/db"
	// "github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
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

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	var id uint
	find_user := db.First(&user, "mobile_number = ? AND devise_token = ?", user.MobileNumber, user.DeviseToken)
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
		find_user := db.Table("users").Where("mobile_number = ? AND devise_token = ?", user.MobileNumber, user.DeviseToken).Select("id").Row()
		find_user.Scan(&id)
		check_session_exists_for_mobile_number := db.First(&session, session.User.MobileNumber)
		if check_session_exists_for_mobile_number.RecordNotFound() == true {
			session = models.Session{UserId: id, DeviseToken: user.DeviseToken}
			db.Create(&session)
			device = models.Device{UserId: id, SessionId: session.ID}
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

// func (s *sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {

// 	var u models.User
// 	vars := mux.Vars(req)
// 	devise_token := vars["devise_token"]
// 	u.DeviseToken = devise_token

// 	response, error, user := session(u, false, true)

// 	if error == true {
// 		log.Printf("response: %v \n", response)
// 		b, err := json.Marshal(models.ErrorMessage{
// 			Success: "false",
// 			Error:   response,
// 		})
// 		if err != nil {
// 			panic(err)
// 		}
// 		rw.Header().Set("Content-Type", "application/json")
// 		rw.Write(b)
// 		goto end
// 	} else {
// 		log.Printf("response: %v \n", response)
// 		b, err := json.Marshal(models.Message{
// 			User:    user,
// 			Success: "true",
// 			Message: response,
// 		})

// 		if err != nil {
// 			panic(err)
// 		}
// 		rw.Header().Set("Content-Type", "application/json")
// 		rw.Write(b)
// 	}
// end:
// }
