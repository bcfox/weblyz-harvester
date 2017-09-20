package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/weblyz-harvester/api"
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
	user := os.Getenv("HARVESTER_DB_USER")
	host := os.Getenv("HARVESTER_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("HARVESTER_DB_PORT"))
	dbname := os.Getenv("HARVESTER_DB_NAME")
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
	endpoints := &api.Endpoints{DB: db}
	router := mux.NewRouter()
	router.HandleFunc("/feeds", endpoints.GetFeeds).Methods("GET")
	http.ListenAndServe(":8080", router)
}
