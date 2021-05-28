package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "championsclub123"
	dbname   = "postgres"
)

func loadSQLFile(db *sql.DB, sqlFile string) error {
	fmt.Println("Seeding " + sqlFile + "...")

	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func executeSeed() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Ready to seed!")

	files, err := ioutil.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		loadSQLFile(db, "data/"+f.Name())
	}
}

func main() {
	executeSeed()
}
