package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies = Movies{
	Movie{"PELI 1", 2020, "Yo1"},
	Movie{"PELI 2", 1999, "Yo2"},
	Movie{"PELI 3", 2022, "Yo3"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HOME PAGE")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func MovieById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "Pelicula por id: %s", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	
	var movieData Movie
	err:=decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	
	log.Println(movieData)

	movies = append(movies, movieData)
}
