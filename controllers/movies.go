package controllers

import (
	"log"
	"net/http"

	"../functions"
	"github.com/gorilla/mux"
)

// MoviesList controller get all documents from db
func MoviesList(w http.ResponseWriter, r *http.Request) {
	list, err := functions.MoviesFindAll()
	if err != nil {
		functions.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	functions.ResponseWithJSON(w, http.StatusOK, list)
}

// MovieDetails returns one movie from database
func MovieDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := functions.MoviesFindOne(params["id"])
	if err != nil {
		functions.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		log.Fatal(err)
	}
	if movie.Name == "" {
		functions.ResponseWithError(w, http.StatusNotFound, "Not found this movie in out collection")
	}
	functions.ResponseWithJSON(w, http.StatusOK, movie)
}

// MovieNew store a new movie in MongoDB
func MovieNew(w http.ResponseWriter, r *http.Request) {
	movie, err := functions.MoviesCreate(r)
	if err != nil {
		functions.ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}
	functions.ResponseWithJSON(w, http.StatusOK, movie)
}

// MovieUpdate updates mongoDB doc content
func MovieUpdate(w http.ResponseWriter, r *http.Request) {}

// MovieDelete remove one doc from MongoDB
func MovieDelete(w http.ResponseWriter, r *http.Request) {}
