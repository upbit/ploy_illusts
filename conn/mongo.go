package conn

import (
	"os"
	"sync"

	"github.com/spf13/viper"
	log "github.com/upbit/ploy_illusts/logger"
	"gopkg.in/mgo.v2"
)

var (
	onceMongo sync.Once
	mdb       *mgo.Database
)

// GetMongoDB GetMongoDB
func GetMongoDB() *mgo.Database {
	onceMongo.Do(func() {
		host := viper.GetString("mongo.host")
		dbName := viper.GetString("mongo.db")
		session, err := mgo.Dial(host)
		if err != nil {
			log.Errorf("mgo.Dial(%s) error: %s", host, err.Error())
			os.Exit(1)
		}

		session.SetMode(mgo.Monotonic, true)
		mdb = session.DB(dbName)
	})
	return mdb
}
