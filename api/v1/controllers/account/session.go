package account

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
}

func (s *sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
}
