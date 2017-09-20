package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/weblyz-harvester/dao"
)

type Endpoints struct {
	DB *sql.DB
}

func (endpoints *Endpoints) GetFeeds(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	feeds := dao.GetAllFeeds(endpoints.DB)
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
