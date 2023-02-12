package main

type Movie struct {
	Name     string `json:"name" bson:"name,omitempty"`
	Year     int    `json:"year" bson:"year,omitempty"`
	Director string `json:"director" bson:"director,omitempty"`
}

type Movies []Movie

type MoviesDoc []interface{}
