package main

import (
	controllers "golangMongoDB/controller"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r:= httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	MongoDB_URI := os.Getenv("MONGO_URI")
	s, err := mgo.Dial(MongoDB_URI)
	if err != nil {
		panic(err)
	}
	return s
}