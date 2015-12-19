package users

import (
	"database/sql"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/kedarnag13/Marga/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
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
	var no_of_users int
	for get_all_issues.Next() {
		err := get_all_issues.Scan(&name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
		if err != nil {
			log.Fatal(err)
		}
		issue_det := models.IssueDetails{name, issue_type, description, latitude, longitude, image, status, address, user_id}
		i.Issue_Details = append(i.Issue_Details, issue_det)
		no_of_users++
		flag = 0
	}
	if flag == 0 {
		b, err := json.Marshal(models.IssueList{
			Success:       "true",
			No_Of_Users:   no_of_users,
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
			issue_res, err := prepare_insert_issue.Exec(id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, id)
			if err != nil || issue_res == nil {
				log.Fatal(err)
			}
			issue := models.Issue{id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, id}
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
