package mdb

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "github.com/smartystreets/goconvey/convey"
)

type Task struct {
	// TaskId     string "bson:`_id"
	Id         int32  "bson:`id"
	Type       string "bson:`type"
	Data       string "bson:`data"
	LastCreate int32  "bson:`last_create_ts"
	Modify     int32  "bson:`modify"
}

func Test_accepting_new_client_callback(t *testing.T) {
	db := getDB("localhost:27017", "pixiv")
	Convey("DB should NOT be nil", t, func() {
		So(db, ShouldNotBeNil)
	})

	coll := db.C("tasks")
	Convey("Collection should NOT be nil", t, func() {
		So(coll, ShouldNotBeNil)
	})

	task := Task{}
	err := coll.Find(bson.M{"id": 853087}).One(&task)
	Convey("Find should be successed", t, func() {
		So(err, ShouldBeNil)
	})
	Convey("Task.id should be equal", t, func() {
		So(task.Id, ShouldEqual, 853087)
	})
	Convey("Task.data should NOT be empty", t, func() {
		So(task.Data, ShouldNotBeEmpty)
	})
}
