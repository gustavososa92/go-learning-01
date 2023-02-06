package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", Contact)

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)

}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo Web")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pantalla Contacto")
}
