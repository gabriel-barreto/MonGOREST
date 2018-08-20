package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"./config"
	"./controllers"
)

var (
	// ROOTROUTER is API router
	ROOTROUTER = mux.NewRouter()
	// APIROUTER recieve pathPrefix /api
	APIROUTER = ROOTROUTER.PathPrefix("/api").Subrouter()
	// MOVIESROUTER is router to movies endpoints
	MOVIESROUTER = APIROUTER.PathPrefix("/movies").Subrouter()
)

func main() {
	// ==> Inject routes
	// => Movies
	MOVIESROUTER.HandleFunc("", controllers.MoviesList).Methods("GET")
	MOVIESROUTER.HandleFunc("", controllers.MovieNew).Methods("POST")
	MOVIESROUTER.HandleFunc("/{id}", controllers.MovieDetails).Methods("GET")

	// ==> Logging API init
	fmt.Println("API in Online!")
	fmt.Println(" ===> PORT:", config.PORT)

	// ==> Start server
	http.Handle("/", ROOTROUTER)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
