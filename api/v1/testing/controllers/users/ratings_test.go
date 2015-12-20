package users

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/users"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ratingsController struct{}

var Ratings ratingsController

func TestCreateIssue(t *testing.T) {
	inputJson := strings.NewReader(`{"SenderId":1,"ReciverId":2}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/user_points", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/user_points", users.Ratings.Create).Methods("POST")
	m.ServeHTTP(w, r)
}
