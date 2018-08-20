package functions

import (
	"log"

	"../config"
	mgo "gopkg.in/mgo.v2"
)

// Connect returns a mongo DB
func Connect() *mgo.Database {
	session, err := mgo.Dial(config.URI)
	if err != nil {
		log.Fatal(err.Error())
	}
	return session.DB(config.DB)
}
