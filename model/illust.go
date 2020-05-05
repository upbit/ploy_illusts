package model

import (
	"time"

	"github.com/upbit/ploy_illusts/conn"
	"gopkg.in/mgo.v2/bson"
)

const (
	illustCollection = "illusts"
)

// Illust Illust
type Illust struct {
	ID        int    `bson:"_id"`
	IllustID  int    `bson:"illust_id"`
	Title     string `bson:"title"`
	Type      string `bson:"type"`
	ImageUrls struct {
		SquareMedium string `bson:"square_medium"`
		Medium       string `bson:"medium"`
		Large        string `bson:"large"`
	} `bson:"image_urls"`
	Caption string `bson:"caption"`
	User    User   `bson:"user"`
	Tags    []struct {
		Name           string      `bson:"name"`
		TranslatedName interface{} `bson:"translated_name"`
	} `bson:"tags"`
	CreateDate     time.Time `bson:"create_date"`
	PageCount      int       `bson:"page_count"`
	Width          int       `bson:"width"`
	Height         int       `bson:"height"`
	MetaSinglePage struct {
		OriginalImageURL string `bson:"original_image_url"`
	} `bson:"meta_single_page"`
	TotalView      int `bson:"total_view"`
	TotalBookmarks int `bson:"total_bookmarks"`
	UserID         int `bson:"user_id"`
	TotalComments  int `bson:"total_comments"`
	CreateDateTs   int `bson:"create_date_ts"`
}

// User User
type User struct {
	ID               int    `bson:"id"`
	Name             string `bson:"name"`
	Account          string `bson:"account"`
	ProfileImageUrls struct {
		Medium string `bson:"medium"`
	} `bson:"profile_image_urls"`
}

// Illusts list
type Illusts []Illust

// IllustInfo model function
func IllustInfo(illustID int) (Illust, error) {
	mdb := conn.GetMongoDB()
	var illust Illust
	err := mdb.C(illustCollection).Find(bson.M{"_id": &illustID}).One(&illust)
	return illust, err
}
