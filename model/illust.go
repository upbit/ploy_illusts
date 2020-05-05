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
	ID        int    `bson:"_id" json:"id"`
	IllustID  int    `bson:"illust_id" json:"illust_id"`
	Title     string `bson:"title" json:"title"`
	Type      string `bson:"type" json:"type"`
	ImageUrls struct {
		SquareMedium string `bson:"square_medium" json:"square_medium"`
		Medium       string `bson:"medium" json:"medium"`
		Large        string `bson:"large" json:"large"`
	} `bson:"image_urls" json:"image_urls"`
	Caption string `bson:"caption" json:"caption"`
	User    User   `bson:"user" json:"user"`
	Tags    []struct {
		Name           string      `bson:"name" json:"name"`
		TranslatedName interface{} `bson:"translated_name" json:"translated_name"`
	} `bson:"tags" json:"tags"`
	CreateDate     time.Time `bson:"create_date" json:"create_date"`
	PageCount      int       `bson:"page_count" json:"page_count"`
	Width          int       `bson:"width" json:"width"`
	Height         int       `bson:"height" json:"height"`
	MetaSinglePage struct {
		OriginalImageURL string `bson:"original_image_url" json:"original_image_url"`
	} `bson:"meta_single_page" json:"meta_single_page"`
	TotalView      int `bson:"total_view" json:"total_view"`
	TotalBookmarks int `bson:"total_bookmarks" json:"total_bookmarks"`
	UserID         int `bson:"user_id" json:"user_id"`
	TotalComments  int `bson:"total_comments" json:"total_comments"`
	CreateDateTs   int `bson:"create_date_ts" json:"create_date_ts"`
}

// User User
type User struct {
	ID               int    `bson:"id" json:"id"`
	Name             string `bson:"name" json:"name"`
	Account          string `bson:"account" json:"account"`
	ProfileImageUrls struct {
		Medium string `bson:"medium" json:"medium"`
	} `bson:"profile_image_urls" json:"profile_image_urls"`
}

// GetIllust get one illust by ID
func GetIllust(illustID int) (*Illust, error) {
	mdb := conn.GetMongoDB()
	illust := new(Illust)
	err := mdb.C(illustCollection).Find(bson.M{"_id": &illustID}).One(illust)
	return illust, err
}

// GetIllusts get all illusts
func GetIllusts(page, size int, sortFields []string) ([]*Illust, error) {
	mdb := conn.GetMongoDB()
	illusts := make([]*Illust, 0, size)
	err := mdb.C(illustCollection).Find(bson.M{}).Sort(sortFields...).
		Skip(page * size).Limit(size).All(&illusts)
	return illusts, err
}
