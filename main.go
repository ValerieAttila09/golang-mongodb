package main

import (
	"go-mongo-api/config"
	"go-mongo-api/controllers"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	controllers.SetBookCollection()

	r := mux.NewRouter()

	// CRUD Routes
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	log.Println("Server running on port 5231")
	http.ListenAndServe(":5231", r)
}
