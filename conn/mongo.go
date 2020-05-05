package mongo

import (
	"gopkg.in/mgo.v2"
)

// DB DB
type DB struct {
	db *mgo.Database
}

// NewDB NewDB
func NewDB(host, dbName string) (*DB, error) {
	mdb := &DB{}
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	mdb.db = session.DB(dbName)
	return mdb, nil
}
