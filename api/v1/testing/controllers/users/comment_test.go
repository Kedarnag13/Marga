package users

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/users"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type commentController struct{}

var Comment commentController

func TestCreateComment(t *testing.T) {
	inputJson := strings.NewReader(`{"User_id":1,"Issue_id":55,"Description":"i can help you"}`)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/create_comment", inputJson)
	m := mux.NewRouter()
	m.HandleFunc("/create_comment", users.Comment.Create).Methods("POST")
	m.ServeHTTP(w, r)
}
