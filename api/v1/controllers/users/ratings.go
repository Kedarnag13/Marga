package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/kedarnag13/Marga/api/v1/controllers"
	"github.com/asaskevich/govalidator"
	"github.com/kedarnag13/Marga/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ratingsController struct{}

var Ratings ratingsController

func (r ratingsController) MyPionts(rw http.ResponseWriter, req *http.Request){

	var point Mypoints

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

	sender_existance := controllers.Check_for_user(point.SenderId)
	sender_session_existance := controllers.Check_for_user_session(point.SenderId)
	receiver_existance := controllers.Check_for_user(point.ReciverId)

	 if sender_existance == false {
		b, err := json.Marshal(models.ProfileErrorMessage{
			Success: "false",
			Error:   "User Does not exist",
			})

		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	}
	if sender_session_existance == false {
		b, err := json.Marshal(models.ProfileErrorMessage{
			Success: "false",
			Error:   "Require Login",
			})

		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	}
	if point.Sender_user_id == 0 ||  point.Receiver_user_id == 0 {
		result, err := govalidator.ValidateStruct(point)
		if err != nil || result == false {
			println("error: " + err.Error())
		}
		b, err := json.Marshal(models.ShowErrorMessage{
			Success: "false",
			Error:   err.Error(),
			})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else if receiver_existance == false {
		b, err := json.Marshal(models.ShowErrorMessage{
			Success: "false",
			Error:   "The receiver does not exist",
			})
		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {

		fetch_point, err := db.Query("select my_points from users where id = $1",point.ReciverId)
			if err != nil {
				log.Fatal(err)
			}

			for fetch_point.Next(){
				var existing_points int
				err = fetch_point.Scan(&existing_points)
				if err != nil {
					log.Fatal(err)
				}
				new_points := existing_points + point.Points
				update_point, err := db.Query("insert into user (my_points) values ($1) where id = $2 ", new_points, point.ReciverId)
				if err != nil {
					log.Fatal(err)
				}

				b, err := json.Marshal(models.SuccessMessage{
					Success: "false",
					Message:   "You have rated successfully",
				})

				if err != nil {
					log.Fatal(err)
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.Write(b)
				fmt.Println("You have rated successfully")
			}
	}
	:end
}