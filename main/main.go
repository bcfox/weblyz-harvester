package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "foxrds.cty3szl04mof.us-east-2.rds.amazonaws.com"
	port   = 5432
	user   = "bfox"
	dbname = "tracker"
)

// https://www.calhoun.io/using-postgresql-with-golang/
func main() {
	password := os.Getenv("HARVESTER_DB_PASSWORD")
	fmt.Println("hello world")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	getFeeds(db)
}

func getFeeds(db *sql.DB) {
	rows, err := db.Query("SELECT fd_uuid, fd_url, fd_name, fd_group FROM feeds")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var url string
		var name string
		var group string
		err = rows.Scan(&id, &url, &name, &group)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, url, name, group)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
