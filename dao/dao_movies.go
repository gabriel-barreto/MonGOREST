package dao

import (
	"gopkg.in/mgo.v2"
)

const (
	// COLLECTION set DAO collection name
	COLLECTION = "movies"
)

// Movies is a Mongo movies collection
func Movies(db *mgo.Database) *mgo.Collection {
	return db.C(COLLECTION)
}
