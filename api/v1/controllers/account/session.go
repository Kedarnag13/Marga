package account

type sessionController struct{}

var Session sessionController

func (s sessionController) Create(rw http.ResponseWriter, req *http.Request) {
}

func (s *sessionController) Destroy(rw http.ResponseWriter, req *http.Request) {
}
