package mdb

import (
	mgo "gopkg.in/mgo.v2"
)

func getDB(addr string, db_name string) *mgo.Database {
	session, err := mgo.Dial(addr)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	return session.DB(db_name)
}
