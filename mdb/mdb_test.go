package mdb

import (
	"fmt"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2/bson"
)

const (
	CHECK_ILLUST_ID = 56102119
)

type CMetaSinglePage struct {
	OriginalImageUrl string "bson:`original_image_url"
}

type CMetaPages struct {
	SquareMedium []string "bson:`square_medium"
	Large        []string "bson:`large"
	Medium       []string "bson:`medium"
	Original     []string "bson:`original"
}

type CImageUrl struct {
	SquareMedium string "bson:`square_medium"
	Large        string "bson:`large"
	Medium       string "bson:`medium"
}

type CUser struct {
	// ObjId            string    "bson:`_id"
	Account          string    "bson:`account"
	ProfileImageUrls CImageUrl "bson:`profile_image_urls"
	Id               int32     "bson:`id"
	Name             string    "bson:`name"
}

type CIllust struct {
	// ObjId          bson.ObjectId   "bson:`_id"
	CreateDate     string          "bson:`create_date"
	TotalComments  int32           "bson:`total_comments"
	PageCount      int32           "bson:`page_count"
	Height         int32           "bson:`height"
	Restrict       int32           "bson:`restrict"
	TotalView      int32           "bson:`total_view"
	Tools          []string        "bson:`tools"
	Id             int32           "bson:`id"
	MetaSinglePage CMetaSinglePage "bson:`meta_single_page"
	MetaPages      CMetaPages      "bson:`meta_pages"
	Title          string          "bson:`title"
	Weight         int32           "bson:`weight"
	Type           string          "bson:`type"
	Tags           []string        "bson:`tags"
	ImageUrls      CImageUrl       "bson:`image_urls"
	User           CUser           "bson:`user"
	CreateDateTs   int32           "bson:`create_date_ts"
	SanityLevel    int32           "bson:`sanity_level"
	Caption        string          "bson:`caption"
	TotalBookmarks int32           "bson:`total_bookmarks"
}

func init() {
	viper.SetConfigName("config_test")
	viper.SetConfigType("json")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s", err)
	}
}

func Test_accepting_new_client_callback(t *testing.T) {
	connections := viper.GetString("mongo.conn")
	db := getDB(connections, "pixiv")
	Convey("DB should NOT be nil", t, func() {
		So(db, ShouldNotBeNil)
	})

	coll := db.C("illusts")
	Convey("Collection should NOT be nil", t, func() {
		So(coll, ShouldNotBeNil)
	})

	//illust := CIllust{}
	iter := coll.Find(bson.M{"id": CHECK_ILLUST_ID}).Iter()

	var _raw interface{}
	for iter.Next(&_raw) {
		fmt.Println(_raw)
	}

	// err := coll.Find(bson.M{"id": CHECK_ILLUST_ID}).One(&illust)
	// Convey("Find by id should be successed", t, func() {
	// 	So(err, ShouldBeNil)
	// })
	// Convey("id should be equal", t, func() {
	// 	So(illust.Id, ShouldEqual, CHECK_ILLUST_ID)
	// })
	// Convey("title should NOT be empty", t, func() {
	// 	So(illust.Title, ShouldNotBeEmpty)
	// })

	// Convey("meta_single_page.original_image_url should NOT be empty", t, func() {
	// 	So(illust.MetaSinglePage.OriginalImageUrl, ShouldNotBeEmpty)
	// })

	// Convey("user.name should NOT be empty", t, func() {
	// 	So(illust.User.Name, ShouldNotBeEmpty)
	// })
	// Convey("user.profile_image_urls.medium should NOT be empty", t, func() {
	// 	So(illust.User.ProfileImageUrls.Medium, ShouldNotBeEmpty)
	// })

	// t.Log(illust)
}
