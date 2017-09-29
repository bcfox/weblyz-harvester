package dao

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/weblyz-harvester/model"
)

// FeedDao DAO for feeds articles and RawSyndications
type FeedDao struct {
	DB *sql.DB
}

// https://www.calhoun.io/using-postgresql-with-golang/

// GetAllFeeds obtains feed objects from the database
func (dao *FeedDao) GetAllFeeds() []model.Feed {
	var feeds = make([]model.Feed, 0)
	rows, err := dao.DB.Query("SELECT fd_uuid, fd_url, fd_name, fd_group FROM feeds")
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
		feeds = append(feeds, model.Feed{ID: id, URL: url, Name: name, Group: group})
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return feeds
}

// SaveArticle if article exists it updates it else inserts
func (dao *FeedDao) SaveArticle(modelFeed *model.Feed, article *model.Article, batchID string) {

}

func (dao *FeedDao) articleExists(modelFeed *model.Feed, article *model.Article, batchID string) {
}

func (dao *FeedDao) insertArticle(modelFeed *model.Feed, article *model.Article, batchID string) {
	sqlStatement := `
		INSERT INTO Articles(uuid_generate_v4(),
			art_fd_uuid, ,art_title, art_url, art_summary,
			art_pub_date, art_updated_date, batchId)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := dao.DB.Exec(sqlStatement, modelFeed.ID, article.Title,
		article.URL, article.Summary, article.PubDate, article.UpdatedDate,
		batchID)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (dao *FeedDao) SaveRaw(raw *model.RawSyndication) {
	uuid := uuid.New()
	raw.ID = uuid.String()

	sqlStatement := `
		INSERT INTO RAW_SYNDICATION(rs_uuid, rs_fd_uuid, rs_pulled_date, rs_data,
				rs_content_type, rs_format, rs_batch_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := dao.DB.Exec(sqlStatement,
		raw.ID, raw.FeedID, raw.PulledDate, raw.Data, raw.ContentType,
		raw.Format, raw.BatchID)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
