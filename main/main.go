package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/weblyz-harvester/dao"
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
	endpoints := &Endpoints{db: db}
	router := mux.NewRouter()
	router.HandleFunc("/feeds", endpoints.GetFeeds).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handleFeeds(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	feeds := dao.GetAllFeeds(db)
	for _, feed := range feeds {
		fmt.Println(feed)
	}

	outgoingJSON, error := json.Marshal(feeds)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}
