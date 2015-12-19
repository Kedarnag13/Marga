package users

import (
	"database/sql"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers"
	"github.com/kedarnag13/Marga/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type issueController struct{}

var Issue issueController

func (is issueController) Index(rw http.ResponseWriter, req *http.Request) {

	var i models.IssueList

	flag := 1
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
	get_all_issues, err := db.Query("SELECT name, type, description, latitude, longitude, image, status, address, user_id  FROM issues")
	if err != nil || get_all_issues == nil {
		log.Fatal(err)
	}
	var name string
	var issue_type string
	var description string
	var latitude float64
	var longitude float64
	var image string
	var status bool
	var address string
	var user_id int
	var no_of_issues int
	for get_all_issues.Next() {
		err := get_all_issues.Scan(&name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
		if err != nil {
			log.Fatal(err)
		}
		issue_det := models.IssueDetails{name, issue_type, description, latitude, longitude, image, status, address, user_id}
		i.Issue_Details = append(i.Issue_Details, issue_det)
		no_of_issues++
		flag = 0
	}
	if flag == 0 {
		b, err := json.Marshal(models.IssueList{
			Success:       "true",
			No_Of_Issues:  no_of_issues,
			Issue_Details: i.Issue_Details,
		})
		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto index_end
	}
	if flag == 1 {
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "No Issues",
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
index_end:
}

func (m issueController) My_issues(rw http.ResponseWriter, req *http.Request) {

	var my_issues models.IssueList
	var no_of_issues int

	flag := 1

	vars := mux.Vars(req)
	id := vars["id"]
	issue_id, err := strconv.Atoi(id)

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	get_issues, err := db.Query("select name, type, description, latitude, longitude, image, status, address, user_id from issues where id = $1 ", issue_id)
	if err != nil {
		log.Fatal(err)
	}

	user_session_existance := controllers.Check_for_user_session(issue_id)
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
		goto my_issue_index_end
	}

	if flag == 0 {
		var name string
		var issue_type string
		var description string
		var latitude float64
		var longitude float64
		var image string
		var status bool
		var address string
		var user_id int
		for get_issues.Next() {

			err := get_issues.Scan(&name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
			if err != nil {
				log.Fatal(err)
			}
			issue_det := models.IssueDetails{name, issue_type, description, latitude, longitude, image, status, address, user_id}
			my_issues.Issue_Details = append(my_issues.Issue_Details, issue_det)
			no_of_issues++
			flag = 0
		}
	}

	if flag == 0 {
		b, err := json.Marshal(models.IssueList{
			Success:       "true",
			No_Of_Issues:  no_of_issues,
			Issue_Details: my_issues.Issue_Details,
		})
		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto my_issue_index_end
	}
	if flag == 1 {
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "No Issues",
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
my_issue_index_end:
}

func (m issueController) Get_issues_on_type(rw http.ResponseWriter, req *http.Request) {

	var my_issues models.IssueList
	var no_of_issues int

	flag := 1

	vars := mux.Vars(req)
	issue_type := vars["type"]

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	get_issues, err := db.Query("select name, type, description, latitude, longitude, image, status, address, user_id from issues where type = $1 ", issue_type)
	if err != nil {
		log.Fatal(err)
	}

	if flag == 1 {
		var name string
		var issue_type string
		var description string
		var latitude float64
		var longitude float64
		var image string
		var status bool
		var address string
		var user_id int
		for get_issues.Next() {

			err := get_issues.Scan(&name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
			if err != nil {
				log.Fatal(err)
			}
			issue_det := models.IssueDetails{name, issue_type, description, latitude, longitude, image, status, address, user_id}
			my_issues.Issue_Details = append(my_issues.Issue_Details, issue_det)
			no_of_issues++
			flag = 0
		}
	}

	if flag == 0 {
		b, err := json.Marshal(models.IssueList{
			Success:       "true",
			No_Of_Issues:  no_of_issues,
			Issue_Details: my_issues.Issue_Details,
		})
		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto issue_on_type
	}
	if flag == 1 {
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "No Issues",
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
issue_on_type:
}

func (is issueController) Create(rw http.ResponseWriter, req *http.Request) {

	var i models.Issue

	flag := 1
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &i)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
	fetch_id, err := db.Query("SELECT coalesce(max(id), 0) FROM issues")
	if err != nil {
		log.Fatal(err)
	}

	if flag == 1 {
		if i.Name == "" || i.Type == "" || i.Description == "" || i.Image == "" || i.Status == false || i.Address == "" {
			result, err := govalidator.ValidateStruct(i)
			if err != nil || result == false {
				println("error: " + err.Error())
			}
			flag = 0
			b, err := json.Marshal(models.IssueErrorMessage{
				Success: "false",
				Error:   err.Error(),
			})
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto issue_end
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
			var insert_issue string = "insert into issues (id, name, type, description, latitude, longitude, image, status, address, user_id) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
			prepare_insert_issue, err := db.Prepare(insert_issue)
			if err != nil || prepare_insert_issue == nil {
				log.Fatal(err)
			}
			issue_res, err := prepare_insert_issue.Exec(id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, i.User_id)
			if err != nil || issue_res == nil {
				log.Fatal(err)
			}
			issue := models.Issue{id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, i.User_id}
			b, err := json.Marshal(models.SuccessfulCreateIssue{
				Success: "true",
				Message: "Issue created Successfully!",
				Issue:   issue,
			})

			if err != nil {
				log.Fatal(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
			goto issue_end
		}
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "User does not exist",
		})

		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
issue_end:
}
