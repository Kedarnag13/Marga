package users

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/users"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type issueController struct{}

var Issue issueController

func TestCreateIssue(t *testing.T) {
	inputJson := strings.NewReader(`{"Name":"praveen issue","Type":"postman","Description":"asjdhkasjdhk","Latitude":12.311790,"Longitude":76.652059,"Image":"shsh.png","status":true,
"Address":"kasjdbkasjdhask","User_id":1}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/create_issue", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/create_comment", users.Comment.Create).Methods("POST")
	m.ServeHTTP(w, r)
}

// func TestIndexIssue(t *testing.T) {

// }
