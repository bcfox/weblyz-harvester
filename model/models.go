package model

import (
	"time"
)

// Feed dto
type Feed struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

// RawSyndication holds a document pulled for a feed
type RawSyndication struct {
	ID          string    `json:"id"`
	FeedID      string    `json:"feedId"`
	PulledDate  time.Time `json:"pulledDate"`
	Data        string    `json:"data"`
	ContentType string    `json:"contentType"`
	Format      string    `json:"format"`
	BatchID     string    `json:"batchID"`
}

// Article dto, holds artical info
type Article struct {
	ID          string    `json:"id"`
	FeedID      string    `json:"feedId"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	PubDate     time.Time `json:"pubDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	BatchID     string    `json:"batchId"`
}
