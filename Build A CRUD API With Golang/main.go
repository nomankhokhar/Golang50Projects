package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    "json:id"
	Isbn     string    "json:isbn"
	Title    string    "json:title"
	Director *Director "json:director"
}

type Director struct {
	FirstName string "json:firstname"
	LastName  string "json:lastname"
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", r)
}

func getMovie(){
	fmt.Println("Get Movie")
}

func getMovies(){
	fmt.Println("Get Movies")
}

func createMovie(){
	fmt.Println("Create Movie")
}

func updateMovie(){
	fmt.Println("Update Movie")
}

func deleteMovie(){
	fmt.Println("Delete Movie")
}