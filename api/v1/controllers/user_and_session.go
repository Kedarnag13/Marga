package controllers

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Check_for_user(user_id int) bool {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	user_ids, err := db.Query("SELECT id FROM users where id = $1", user_id)
	if err != nil {
		panic(err)
	}
	flag := 1
	for user_ids.Next() {
		var id int
		err = user_ids.Scan(&id)
		if err != nil {
			panic(err)
		}

		if id == user_id {
			flag = 0
		}
	}
	defer user_ids.Close()
	db.Close()
	if flag == 0 {
		return true
	} else {
		return false
	}
}

func Check_for_user_session(user_id int) bool {
	db, err := sql.Open("postgres", "password=password host=localhost dbname=marga_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	session_user_ids, err := db.Query("SELECT user_id FROM sessions where user_id = $1", user_id)
	if err != nil {
		panic(err)
	}
	flag := 1
	for session_user_ids.Next() {
		var id int
		err = session_user_ids.Scan(&id)
		if err != nil {
			panic(err)
		}

		if id == user_id {
			flag = 0
		}
	}
	defer session_user_ids.Close()
	db.Close()
	if flag == 0 {
		return true
	} else {
		return false
	}
}
