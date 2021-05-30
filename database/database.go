package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "championsclub123"
	dbname   = "postgres"
)

func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
