package users

import (
	"encoding/json"
	"github.com/Kedarnag13/Marga/api/v1/controllers"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/Kedarnag13/Marga/api/v1/config/db"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type issueController struct{}

var Issue issueController

func (is issueController) Index(rw http.ResponseWriter, req *http.Request) {

	var i models.IssueList

	flag := 1

	get_all_issues, err := db.DBCon.Query("SELECT id, name, type, description, latitude, longitude, image, status, address, user_id  FROM issues")
	if err != nil || get_all_issues == nil {
		panic(err)
	}
	defer get_all_issues.Close()
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
			panic(err)
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
			panic(err)
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
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
	index_end:
}

func (m issueController) MyIssues(rw http.ResponseWriter, req *http.Request) {

	var my_issues models.IssueList
	var no_of_issues int

	flag := 0

	vars := mux.Vars(req)
	id := vars["id"]
	user_id, err := strconv.Atoi(id)

	get_issues, err := db.DBCon.Query("select id, name, type, description, latitude, longitude, image, status, address, user_id from issues where user_id = $1 ", user_id)
	if err != nil {
		panic(err)
	}
	defer get_issues.Close()
	user_session_existance := controllers.Check_for_user_session(user_id)
	if user_session_existance == false {
		b, err := json.Marshal(models.ProfileErrorMessage{
			Success: "false",
			Error:   "Require Login",
		})

		if err != nil {
			panic(err)
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
				panic(err)
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
			panic(err)
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
				panic(err)
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

		get_issues, err := db.DBCon.Query("select id, name, type, description, latitude, longitude, image, status, address, user_id from issues where type = $1 ", issue_type)
		if err != nil {
			panic(err)
		}
		defer get_issues.Close()
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
					panic(err)
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
				panic(err)
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
				panic(err)
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

		fetch_id, err := db.DBCon.Query("SELECT coalesce(max(id), 0) FROM issues")
		if err != nil {
			panic(err)
		}
		defer fetch_id.Close()
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
					panic(err)
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
					panic(err)
				}
				id = id + 1

				created_at := time.Now()
				if err != nil {
					panic(err)
				}
				prepare_insert_issue, err := db.DBCon.Prepare("insert into issues (id, name, type, description, latitude, longitude, image, status, address, user_id, created_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
				if err != nil || prepare_insert_issue == nil {
					panic(err)
				}
				issue_res, err := prepare_insert_issue.Exec(id, i.Name, i.Type, i.Description, i.Latitude, i.Longitude, i.Image, i.Status, i.Address, i.User_id, created_at)
				if err != nil || issue_res == nil {
					panic(err)
				}
				defer prepare_insert_issue.Close()
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
				panic(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
		}
		issue_end:
	}

	func (ic issueController) List_wards(rw http.ResponseWriter, req *http.Request) {

		var ward models.WardList
		flag := 1

		get_wards, err := db.DBCon.Query("SELECT id, name, email, devise_token  FROM wards")
		if err != nil || get_wards == nil {
			panic(err)
		}
		defer get_wards.Close()
		var no_of_wards int
		for get_wards.Next() {
			var id int
			var name string
			var email string
			var devise_token string
			err = get_wards.Scan(&id, &name, &email, &devise_token)
			if err != nil {
				panic(err)
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
				panic(err)
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
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
		}
		ward_issue_end:
	}


	func (is issueController) Cluster(rw http.ResponseWriter, req *http.Request) {

		var u models.ClusterIssues
		var issue models.IssueList

		flag := 1

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
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

		for i := 0 ; i < len(u.Issues) ; i++ {
			get_cluster_issues, err := db.DBCon.Query("SELECT id, name, type, description, latitude, longitude, image, status, address, user_id  FROM issues where id = $1",u.Issues[i])
			if err != nil || get_cluster_issues == nil {
				panic(err)
			}
			defer get_cluster_issues.Close()
			for get_cluster_issues.Next() {
				err := get_cluster_issues.Scan(&issue_id, &name, &issue_type, &description, &latitude, &longitude, &image, &status, &address, &user_id)
				if err != nil {
					panic(err)
				}
				issue_det := models.IssueDetails{issue_id, name, issue_type, description, latitude, longitude, image, status, address, user_id}
				issue.Issue_Details = append(issue.Issue_Details, issue_det)
				no_of_issues++
				flag = 0
			}
		}

		if flag == 0 {
			b, err := json.Marshal(models.IssueList{
				Success:       "true",
				No_Of_Issues:  no_of_issues,
				Issue_Details: issue.Issue_Details,
			})
			if err != nil {
				panic(err)
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
				panic(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(b)
		}
		cluster_end:
	}
