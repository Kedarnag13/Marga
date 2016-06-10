package main

import (
	"database/sql"
	"github.com/Kedarnag13/Marga/api/v1/controllers/account"
	"github.com/Kedarnag13/Marga/api/v1/controllers/users"
	"github.com/Kedarnag13/Marga/api/v1/config/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	var err error
	db.DBCon, err = sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.DBCon.Close()
	r := mux.NewRouter()
	// Account Routes
	r.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	r.HandleFunc("/log_in", account.Session.Create).Methods("POST")
	r.HandleFunc("/log_out/{devise_token:([a-zA-Z0-9]+)?}", account.Session.Destroy).Methods("GET")

	// Issue Routes
	r.HandleFunc("/create_issue", users.Issue.Create).Methods("POST")
	r.HandleFunc("/issues", users.Issue.Index).Methods("GET")
	r.HandleFunc("/user/{id:[0-9]+}/issues/{id:[0-9]+}", users.Issue.Index).Methods("GET")
	r.HandleFunc("/issues/{type:[a-z]+}", users.Issue.Get_issues_on_type).Methods("GET")
	// List My Issues
	r.HandleFunc("/user/{id:[0-9]+}/issues", users.Issue.MyIssues).Methods("GET")

	// List Issues in Cluster
	r.HandleFunc("/cluster/issues", users.Issue.Cluster).Methods("POST")

	// Ward Routes
	r.HandleFunc("/wards", users.Issue.List_wards).Methods("GET")

	//Comment
	r.HandleFunc("/create_comment", users.Comment.Create).Methods("POST")
	r.HandleFunc("/comment/issues/{id:[0-9]+}", users.Comment.Index).Methods("GET")

	// Rating Routes
	r.HandleFunc("/user_points", users.Ratings.Create).Methods("POST")

	//Send forgot password message
	r.HandleFunc("/forgot_password", account.ForgotPassword.SendPassword).Methods("POST")
	//Reset Password
	r.HandleFunc("/reset_password", account.ForgotPassword.ResetPassword).Methods("POST")

	http.Handle("/", r)
	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
