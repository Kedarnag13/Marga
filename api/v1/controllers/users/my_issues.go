package users

import (
	"database/sql"
	"encoding/json"
	"github.com/Qwinix/rVidi-Go/api/v1/models"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type myissuesController struct{}

var MyIssues myissuesController

func (m myissuesController) List_users(rw http.ResponseWriter, req *http.Request) {

	var my_issues models.Issue

	vars := mux.Vars(req)
	id := vars["id"]
	tmp, err := strconv.Atoi(id)
	my_issues.UserId = tmp
	type := vars["type"]
	my_issues.IssueType = type

	flag := 1
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	get_issues, err := db.Query("select * from issues where id = $1 AND type =$2", my_issues.UserId, my_issues.IssueType)
	if err != nil {
		log.Fatal(err)
	}

	user_session_existance := controllers.Check_for_user_session(my_issues.UserId)
	if user_session_existance == false {
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

	for get_issues.Next(){
		var id int
		var name string
		var issue_type string
		var description string
		var latitude string
		var longitude string
		var image string
		var status bool
		var address string
		var user_id int

		err := get_issues.Scan(&id,&name,&issue_type,&description,&latitude,&longitude,&image,&status,&address,&user_id)
		if err != nil {
			log.Fatal(err)
		}
		issue_list := models.UserDetails{user_id, firstname, lastname, email, user_thumbnail.String}
		list.User_Details = append(list.User_Details, profile)
		no_of_users++
	}


}