package users

import (
""
)

type issueController struct{}

var Issue issueController

func (i issueController) Create(rw http.ResponseWriter, req *http.Request) {

	var i models.Issue

	flag := 1

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil || db == nil {
		log.Fatal(err)
	}
	check_issues, err := db.Query("SELECT id FROM issues")
	if err != nil || check_issues == nil {
		log.Fatal(err)
	}
	if flag == 1 {
		
	}
}