package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Movie is a model representation from mongo db movie doc
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Cover       string        `bson:"cover" json:"cover"`
	Description string        `bson:"description" json:"description"`
}
