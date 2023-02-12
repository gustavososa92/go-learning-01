package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getSession() *mongo.Client {

	credential := options.Credential{
		AuthSource: "moviesGO",
		Username:   "admin",
		Password:   "password",
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected.")
	return client
}

var moviesCollection = getSession().Database("moviesGO").Collection("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	var movies = MoviesDoc{
		Movie{"PELI 1", 2020, "Yo1"},
		Movie{"PELI 2", 1999, "Yo2"},
		Movie{"PELI 3", 2022, "Yo3"},
	}
	_, err := moviesCollection.InsertMany(context.TODO(), movies)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "HOME PAGE - ADDED MOVIES")
}

func MoviesList(w http.ResponseWriter, r *http.Request) {
	var results []Movie

	cursor, err := moviesCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	defer cursor.Close(context.TODO())

	json.NewEncoder(w).Encode(results)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func MovieById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieId := params["id"]

	fmt.Fprintf(w, "Pelicula por id: %s", movieId)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movieData Movie
	err := decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movieData)

	_, err = moviesCollection.InsertOne(context.TODO(), movieData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(movieData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
