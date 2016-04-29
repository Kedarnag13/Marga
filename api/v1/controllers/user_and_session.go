package controllers

import (
	"github.com/Qwinix/rVidi-Go/api/v1/config/db"
	_ "github.com/lib/pq"
)

func Check_for_user(user_id int) bool {

	user_ids, err := db.DBCon.Query("SELECT id FROM users where id = $1", user_id)
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
	if flag == 0 {
		return true
	} else {
		return false
	}
}

func Check_for_user_session(user_id int) bool {

	session_user_ids, err := db.DBCon.Query("SELECT user_id FROM sessions where user_id = $1", user_id)
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
	if flag == 0 {
		return true
	} else {
		return false
	}
}
