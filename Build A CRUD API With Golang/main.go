package main

import (
	"fmt"
	// "log"
	// "encoding/json"
	// "math/rand"
	// "net/http"
	// "github.com/gorilla/mux"
	// "strconv"
)

type Movie struct{
	ID string "json:id"
	Isbn string "json:isbn"
	Title string "json:title"
	Director *Director "json:director"
}

func main() {
	fmt.Println("Hello World")
}