package users

import (
	"database/sql"
	"encoding/json"
	"github.com/kedarnag13/Marga/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
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
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
	get_issues, err := db.Query("SELECT id from issues where id = $1 AND status = true", c.Issue_id)
	if err != nil || get_issues == nil {
		log.Fatal(err)
	}

	for get_issues.Next() {
		flag = 0
		var issue_id int
		err := get_issues.Scan(&issue_id)
		if err != nil {
			log.Fatal(err)
		}
		var insert_comment string = "insert into comments(user_id,issue_id,description) values ($1,$2,$3)"
		prepare_comments, err := db.Prepare(insert_comment)
		if err != nil {
			log.Fatal(err)
		}
		res, err := prepare_comments.Exec(c.User_id, c.Issue_id, c.Description)
		if err != nil || res == nil {
			log.Fatal(err)
		}
	}
	if flag == 0 {
		b, err := json.Marshal(models.SuccessCommentMessage{
			Success: "true",
			Message: "Comment added successfully",
		})

		if err != nil {
			log.Fatal(err)
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
			log.Fatal(err)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(b)
	}
end:
}
