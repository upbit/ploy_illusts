package model

import (
	"testing"

	"github.com/spf13/viper"

	. "github.com/smartystreets/goconvey/convey"
)

func reloadConfigs(t *testing.T) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		t.Fatalf("Fatal error config file: %s", err)
	}
}

func TestGetIllust(t *testing.T) {
	reloadConfigs(t)

	Convey("It should return Illust as expect", t, func() {
		illust, err := GetIllust(5663)
		So(err, ShouldBeNil)

		t.Log(illust)

		So(illust.ID, ShouldEqual, 5663)
		So(illust.Type, ShouldEqual, "illust")
	})
}

func TestGetIllusts(t *testing.T) {
	reloadConfigs(t)

	Convey("It should return Illusts by page,size", t, func() {
		illusts, err := GetIllusts(0, 20, []string{"_id"})
		So(err, ShouldBeNil)
		So(illusts, ShouldHaveLength, 20)

		So(illusts[0].ID, ShouldEqual, 5663)
		So(illusts[1].ID, ShouldEqual, 5664)
		So(illusts[2].ID, ShouldEqual, 5665)
	})
}
