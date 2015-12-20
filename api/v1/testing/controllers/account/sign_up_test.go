package account

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/account"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type registrationController struct{}

var Registration registrationController

func TestCreateUser(t *testing.T) {
	inputJson := strings.NewReader(`{"Name":"praveen","Username":"praveenmenon","Email":"praveen@yopmail.com","Mobile_number":"9916854301","Latitude":34.56,"Password":"password","Password_confirmation":"password","City":"Mysore","Devise_token":"ksajfhkajsfksajdfkjsghsjkdfhskj","Ward_id":2,"Type":"user"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/sign_up", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/sign_up", account.Registration.Create).Methods("POST")
	m.ServeHTTP(w, r)
}
