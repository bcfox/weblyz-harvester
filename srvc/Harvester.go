package srvc

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"encoding/json"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
	"github.com/weblyz-harvester/dao"
	"github.com/weblyz-harvester/model"
)

type Harvester struct {
	Dao             *dao.FeedDao
	cachedFeeds     []model.Feed
	runPeriodically bool
}

func (srvc *Harvester) HarvestAllPeriodically(secondsPerPeriod int, secondsForPeriod int) {
	srvc.runPeriodically = true
	for srvc.runPeriodically {
		srvc.HarvestAll(secondsForPeriod)
		time.Sleep(time.Duration(secondsPerPeriod) * time.Second)
	}
}

func (srvc *Harvester) HarvestAll(secondsForPeriod int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	batchID := uuid.New().String()
	log.Printf("Starting batch %s\n", batchID)

	if len(srvc.cachedFeeds) < 1 {
		srvc.cachedFeeds = srvc.Dao.GetAllFeeds()
	}

	var wg sync.WaitGroup

	for _, feed := range srvc.cachedFeeds {
		wg.Add(1)
		go func(feed model.Feed, batchID string) {
			defer wg.Done()
			time.Sleep(time.Duration(float32(secondsForPeriod)*r.Float32()) * time.Second)
			log.Printf("  Harvesting batch %s feed %s (%s)\n", batchID, feed.ID, feed.URL)
			srvc.Harvest(&feed, batchID)
		}(feed, batchID)
	}
	wg.Wait()
}

// Harvest pull feed, save the raw then save the articles
func (srvc *Harvester) Harvest(feed *model.Feed, batchID string) *gofeed.Feed {
	fp := gofeed.NewParser()
	parsedFeed, _ := fp.ParseURL(feed.URL)
	data, error := json.Marshal(parsedFeed)
	raw := &model.RawSyndication{
		ID: "", FeedID: feed.ID, PulledDate: time.Now().UTC(),
		Data: string(data), ContentType: "application/json", Format: "gofeed", BatchID: batchID}

	if error != nil {
		log.Println(error.Error())
	}
	srvc.Dao.SaveRaw(raw)
	//articles := srvc.toArticles(feed.ID, batchID, parsedFeed)

	return parsedFeed
}

func (srvc *Harvester) toArticles(feedID string, batchID string, gofeed *gofeed.Feed) (articles []*model.Article) {
	for _, feed := range gofeed.Items {
		article := &model.Article{
			ID: "", FeedID: feedID,
			Title: feed.Title, URL: feed.Link, Summary: feed.Description, Content: feed.Content,
			PubDate: *feed.PublishedParsed, UpdatedDate: *feed.UpdatedParsed, BatchID: batchID}

		articles = append(articles, article)
	}
	return articles
}
