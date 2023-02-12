package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	response(w, http.StatusOK, results)
}

func MovieById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieId := params["id"]

	oid, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	result := Movie{}

	filter := bson.D{{Key: "_id", Value: oid}}
	err = moviesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	response(w, http.StatusOK, result)
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

	response(w, http.StatusOK, movieData)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movieId := params["id"]

	oid, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var movieData Movie
	err = decoder.Decode(&movieData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movieData)

	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.M{"$set": movieData}
	_, err = moviesCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	response(w, http.StatusOK, movieData)
}

func response(w http.ResponseWriter, status int, result interface{}) {
	json.NewEncoder(w).Encode(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
}
