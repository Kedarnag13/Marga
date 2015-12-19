package main

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/account"
	"github.com/kedarnag13/Marga/api/v1/controllers/users"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Account Routes
	r.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	r.HandleFunc("/log_in", account.Session.Create).Methods("POST")
	r.HandleFunc("/log_out/{devise_token:([a-zA-Z0-9]+)?}", account.Session.Destroy).Methods("GET")

	// Issue Routes
	r.HandleFunc("/create_issue", users.Issue.Create).Methods("POST")
	r.HandleFunc("/point_count", users.Ratings.Create).Methods("POST")

	http.Handle("/", r)
	// HTTP Listening Port
	log.Println("main : Started : Listening on: http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
