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

type registrationController struct{}

var Registration registrationController

func TestCreateUser(t *testing.T) {
	var err error
	db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	inputJson := strings.NewReader(`{"name":"Tattoo","username":"Inspiration","email":"tattoo@example.com","password":"password","password_confirmation":"password","city":"mysore","mobile_number":"7022665448","latitude":12345,"longitude":12345,"type":"user","devise_token":"039d51057a2c6125ba53fe6d90daee31837fbc76145dad6186f036cf1d2"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/sign_up", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	m.ServeHTTP(w, r)
}
