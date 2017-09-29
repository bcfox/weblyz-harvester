package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/weblyz-harvester/dao"
	"github.com/weblyz-harvester/srvc"
)

// Endpoints holds model data for the the endpoints
type Endpoints struct {
	Dao       *dao.FeedDao
	Harvester *srvc.Harvester
}

// GetFeeds retrieves all feeds from the database and returns them
func (endpoints *Endpoints) GetFeeds(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	feeds := endpoints.Dao.GetAllFeeds()
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

// ParseAllFeeds parse the feeds
func (endpoints *Endpoints) ParseAllFeeds(res http.ResponseWriter, req *http.Request) {
	feeds := endpoints.Dao.GetAllFeeds()

	for _, feed := range feeds {
		endpoints.Harvester.Harvest(&feed, uuid.New().String())
	}
}
