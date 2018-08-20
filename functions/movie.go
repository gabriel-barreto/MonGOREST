package functions

import (
	"encoding/json"
	"net/http"

	DAO "../dao"
	Models "../models"
	"gopkg.in/mgo.v2/bson"
)

// MoviesFindAll returns all documents from movies collection
func MoviesFindAll() ([]Models.Movie, error) {
	// ==> Getting DAO
	dao := DAO.Movies(Connect())

	var list []Models.Movie
	err := dao.Find(bson.M{}).All(&list)
	return list, err
}

// MoviesFindOne is used to get one document from mongodb using
// DocID that was passed as an argument to this function
func MoviesFindOne(id string) (Models.Movie, error) {
	dao := DAO.Movies(Connect())

	var movie Models.Movie
	err := dao.FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// MoviesCreate receive one movie object like Movie Model
// and register this in mongoDB
func MoviesCreate(req *http.Request) (Models.Movie, error) {
	defer req.Body.Close()

	var payload Models.Movie
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		return payload, err
	}
	payload.ID = bson.NewObjectId()

	dao := DAO.Movies(Connect())
	err = dao.Insert(&payload)
	return payload, err
}

// MoviesUpdate change the data store in mongoDB to
// passed new movie data
func MoviesUpdate(req *http.Request) (Models.Movie, error) {
	defer req.Body.Close()

	var data Models.Movie
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	dao := DAO.Movies(Connect())
	err = dao.UpdateId(data.ID, &data)
	return data, err
}

// MoviesRemove delete one doc from mongoDB
func MoviesRemove(id string) (Models.Movie, error) {
	dao := DAO.Movies(Connect())
	movie, err := MoviesFindOne(id)
	if err != nil {
		return movie, err
	}
	err = dao.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return movie, err
}
