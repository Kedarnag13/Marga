package account

import (
	"database/sql"
"github.com/Kedarnag13/Marga/api/v1/controllers/account"
"github.com/Kedarnag13/Marga/api/v1/config/db"
"net/http"
"net/http/httptest"
"strings"
"github.com/gorilla/mux"
"testing"
)

type sessionController struct{}

var Session sessionController

func TestCreateSession(t *testing.T) {
	var err error
	db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	inputJson := strings.NewReader(`{"password":"password","mobile_number":"7022665448","devise_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/log_in", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/log_in", account.Session.Create).Methods("POST")
	m.ServeHTTP(w, r)
}


func TestDestroySession(t *testing.T) {
	var err error
	db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
  inputURL := strings.NewReader(`{"devise_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/log_out", inputURL)
  account.Session.Destroy(w, r)
}
