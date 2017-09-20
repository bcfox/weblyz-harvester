package dao

import (
	"database/sql"
)

// Feed dto
type Feed struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

// GetAllFeeds obtains feed objects from the database
func GetAllFeeds(db *sql.DB) []Feed {
	var feeds = make([]Feed, 0)
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
		feeds = append(feeds, Feed{ID: id, URL: url, Name: name, Group: group})
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return feeds
}
