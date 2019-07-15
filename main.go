package main

import (
	"log"
	"net/http"

	"github.com/alex-bezverkhniy/gonotes/repositories"
	"github.com/alex-bezverkhniy/gonotes/routers"

	"github.com/alex-bezverkhniy/gonotes/controllers"
)

func main() {
	noteController := controllers.NewNoteController(repositories.NewNoteRepository("data.json"))

	mux := http.NewServeMux()

	routers.CreateRoutes(mux, noteController)

	log.Println("Go Notes Rest API v1.0")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
