package users

import (
	"encoding/json"
	"github.com/Kedarnag13/Marga/api/v1/models"
	"github.com/Kedarnag13/Marga/api/v1/config/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

type commentController struct{}

var Comment commentController

func (is commentController) Create(rw http.ResponseWriter, req *http.Request) {

	var c models.Comment

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		panic(err)
	}

	flag := 1

	get_issues, err := db.DBCon.Query("SELECT id from issues where id = $1 AND status = true", c.Issue_id)
	if err != nil || get_issues == nil {
		panic(err)
	}
	defer get_issues.Close()
	for get_issues.Next() {
		flag = 0
		var issue_id int
		err := get_issues.Scan(&issue_id)
		if err != nil {
			panic(err)
		}
		if issue_id == c.Issue_id {

			prepare_comments, err := db.DBCon.Prepare("insert into comments(description,user_id,issue_id) values ($1,$2,$3)")
			if err != nil {
				panic(err)
			}
			res, err := prepare_comments.Exec(c.Description, c.User_id, c.Issue_id)
			if err != nil || res == nil {
				panic(err)
			}
			defer prepare_comments.Close()
		}
	}
	if flag == 0 {
		b, err := json.Marshal(models.SuccessCommentMessage{
			Success: "true",
			Message: "Comment added successfully",
		})

		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto end
	} else {
		b, err := json.Marshal(models.ErrorMessage{
			Success: "false",
			Error:   "The issue is closed",
		})
		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
end:
}

func (is commentController) Index(rw http.ResponseWriter, req *http.Request) {

	var c models.CommentList
	vars := mux.Vars(req)
	issue_id := vars["id"]
	Issue_id := issue_id

	flag := 1

	get_comments, err := db.DBCon.Query("SELECT description, user_id from comments where issue_id = $1", Issue_id)
	if err != nil || get_comments == nil {
		panic(err)
	}
	defer get_comments.Close()
	var no_of_comment int
	for get_comments.Next() {
		var comment_message string
		var user_id int
		var name string
		err := get_comments.Scan(&comment_message, &user_id)
		if err != nil {
			panic(err)
		}

		get_user_details, err := db.DBCon.Query("SELECT name from users where id= $1", user_id)
		if err != nil || get_user_details == nil {
			panic(err)
		}
		defer get_user_details.Close()
		for get_user_details.Next() {
			err := get_user_details.Scan(&name)
			if err != nil {
				panic(err)
			}
		}

		comment := models.CommentDetails{comment_message, user_id, name}
		c.Comment_details = append(c.Comment_details, comment)
		no_of_comment++
		flag = 0
	}
	defer get_comments.Close()
	if flag == 0 {
		b, err := json.Marshal(models.CommentList{
			Success:         "true",
			No_of_comments:  no_of_comment,
			Comment_details: c.Comment_details,
		})
		if err != nil {
			panic(err)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
		goto comment_end
	}
	if flag == 1 {
		b, err := json.Marshal(models.IssueErrorMessage{
			Success: "false",
			Error:   "No Comments",
		})

		if err != nil {
			panic(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
comment_end:
}
