package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/weblyz-harvester/api"
	"github.com/weblyz-harvester/dao"
	"github.com/weblyz-harvester/srvc"
)

// https://www.calhoun.io/using-postgresql-with-golang/
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	password := os.Getenv("HARVESTER_DB_PASSWORD")
	user := os.Getenv("HARVESTER_DB_USER")
	host := os.Getenv("HARVESTER_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("HARVESTER_DB_PORT"))
	period, err := strconv.Atoi(os.Getenv("HARVESTER_PERIOD"))
	slop, err := strconv.Atoi(os.Getenv("HARVESTER_PERIOD_SLOP"))
	dbname := os.Getenv("HARVESTER_DB_NAME")
	log.Println("hello world")
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
	log.Println("Successfully connected, opening server")
	dao := &dao.FeedDao{DB: db}
	harvester := &srvc.Harvester{Dao: dao}
	endpoints := &api.Endpoints{Harvester: harvester, Dao: dao}
	router := mux.NewRouter()
	router.HandleFunc("/feeds", endpoints.GetFeeds).Methods("GET")
	router.HandleFunc("/feeds/parseAll", endpoints.ParseAllFeeds).Methods("POST")
	go harvester.HarvestAllPeriodically(period, slop)
	http.ListenAndServe(":8080", router)
}
