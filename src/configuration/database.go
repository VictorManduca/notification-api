package configuration

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	dbUrl := Env().DB_URL
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
