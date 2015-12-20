package account

import (
	"github.com/gorilla/mux"
	"github.com/kedarnag13/Marga/api/v1/controllers/account"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDestory(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/log_out/jasdhkjashdiuqwkajbsjdh", nil)
	m := mux.NewRouter()
	r.HandleFunc("/log_out/{devise_token:([a-zA-Z0-9]+)?}", account.Session.Destroy).Methods("GET")
	m.ServeHTTP(w, r)
}
