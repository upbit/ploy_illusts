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

func TestIllustInfo(t *testing.T) {
	reloadConfigs(t)

	Convey("It should return Illust as expect", t, func() {
		illust, err := IllustInfo(5663)
		So(err, ShouldBeNil)

		t.Log(illust)

		So(illust.ID, ShouldEqual, 5663)
		So(illust.Type, ShouldEqual, "illust")
	})
}
