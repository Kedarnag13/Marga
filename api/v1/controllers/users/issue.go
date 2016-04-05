package users

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
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
	get_all_issues, err := db.Query("SELECT id, name, type, description, latitude, longitude, image, status, address, user_id  FROM issues")
	if err != nil || get_all_issues == nil {
		log.Fatal(err)
	}
	var issue_id int
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
		err := get_all_issues.Scan(&issue_id, &name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
		if err != nil {
			log.Fatal(err)
		}
		issue_det := models.IssueDetails{issue_id, name, issue_type, description, latitude, longitude, image, status, address, user_id}
		i.Issue_Details = append(i.Issue_Details, issue_det)
		no_of_issues++
		flag = 0
	}
	defer get_all_issues.Close()
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
db.Close()
}

func (m issueController) MyIssues(rw http.ResponseWriter, req *http.Request) {

	var my_issues models.IssueList
	var no_of_issues int

	flag := 0

	vars := mux.Vars(req)
	id := vars["id"]
	user_id, err := strconv.Atoi(id)

	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	get_issues, err := db.Query("select id, name, type, description, latitude, longitude, image, status, address, user_id from issues where user_id = $1 ", user_id)
	if err != nil {
		log.Fatal(err)
	}

	user_session_existance := controllers.Check_for_user_session(user_id)
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
		var issue_id int
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

			err := get_issues.Scan(&issue_id, &name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
			if err != nil {
				log.Fatal(err)
			}
			issue_det := models.IssueDetails{issue_id, name, issue_type, description, latitude, longitude, image, status, address, user_id}
			my_issues.Issue_Details = append(my_issues.Issue_Details, issue_det)
			no_of_issues++
			flag = 0
		}
	}
	defer get_issues.Close()

	if flag == 0 && no_of_issues > 0 {
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
	} else {
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
db.Close()
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

	get_issues, err := db.Query("select id, name, type, description, latitude, longitude, image, status, address, user_id from issues where type = $1 ", issue_type)
	if err != nil {
		log.Fatal(err)
	}

	if flag == 1 {
		var issue_id int
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

			err := get_issues.Scan(&issue_id, &name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
			if err != nil {
				log.Fatal(err)
			}
			issue_det := models.IssueDetails{issue_id, name, issue_type, description, latitude, longitude, image, status, address, user_id}
			my_issues.Issue_Details = append(my_issues.Issue_Details, issue_det)
			no_of_issues++
			flag = 0
		}
	}
	defer get_issues.Close()

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
db.Close()
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

			created_at := time.Now()
			if err != nil {
				log.Fatal(err)
			}
			var insert_issue string = "insert into issues (id, name, type, description, latitude, longitude, image, status, address, user_id, created_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
			prepare_insert_issue, err := db.Prepare(insert_issue)
			if err != nil || prepare_insert_issue == nil {
				log.Fatal(err)
			}
			issue_res, err := prepare_insert_issue.Exec(id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, i.User_id, created_at)
			if err != nil || issue_res == nil {
				log.Fatal(err)
			}
			issue := models.Issue{id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, i.User_id, i.Corporator_id, created_at}

			b, err := json.Marshal(models.SuccessfulCreateIssue{
				Success: "true",
				Message: "Issue created Successfully!",
				Issue:   issue,
			})
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
db.Close()
}

func (ic issueController) List_wards(rw http.ResponseWriter, req *http.Request) {

	var ward models.WardList
	flag := 1
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
	get_wards, err := db.Query("SELECT id, name, email, devise_token  FROM wards")
	if err != nil || get_wards == nil {
		log.Fatal(err)
	}
	var no_of_wards int
	for get_wards.Next() {
		var id int
		var name string
		var email string
		var devise_token string
		err = get_wards.Scan(&id, &name, &email, &devise_token)
		if err != nil {
			log.Fatal(err)
		}
		ward_det := models.WardDetails{id, name, email, devise_token}
		ward.Ward_Details = append(ward.Ward_Details, ward_det)
		no_of_wards++
		flag = 0
	}
	if flag == 0 {
		b, err := json.Marshal(models.WardList{
			Success:      "true",
			No_Of_Wards:  no_of_wards,
			Ward_Details: ward.Ward_Details,
		})
		if err != nil {
			log.Fatal(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto ward_issue_end
	}
	if flag == 1 {
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "No Wards",
		})

		if err != nil {
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
ward_issue_end:
db.Close()
}


func (is issueController) Cluster(rw http.ResponseWriter, req *http.Request) {

	var u models.ClusterIssues
	var i models.IssueList

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


	get_cluster_issues, err := db.Query("SELECT id, name, type, description, latitude, longitude, image, status, address, user_id  FROM issues where id IN (?)",u.Issues)
	if err != nil || get_cluster_issues == nil {
		log.Fatal(err)
	}
	defer get_cluster_issues.Close()

	var issue_id int
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
	for get_cluster_issues.Next() {
		err := get_cluster_issues.Scan(&issue_id, &name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
		if err != nil {
			log.Fatal(err)
		}
		issue_det := models.IssueDetails{issue_id, name, issue_type, description, latitude, longitude, image, status, address, user_id}
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
		goto cluster_end
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
	cluster_end:
	db.Close()
}
