package config

import (
	"database/sql"
	"github.com/Kedarnag13/Marga/api/v1/config/db"
	"log"
	"os"
)

func Setup_env(go_env string) (env string) {

	var err error

	switch go_env {
	case "it":
		os.Setenv("it", "it")
		log.Println(os.Getenv("it"))
		db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=marga_it sslmode=disable")
		if err != nil {
			panic(err)
		}
		return os.Getenv("it")
	default:
		os.Setenv("development", "development")
		log.Println(os.Getenv("development"))
		db.DBCon, err = sql.Open("postgres", "user=postgres password=password dbname=marga_development sslmode=disable")
		if err != nil {
			panic(err)
		}
		return os.Getenv("development")
	}
}
