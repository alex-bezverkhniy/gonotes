package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alex-bezverkhniy/gonotes/repositories"

	"github.com/alex-bezverkhniy/gonotes/controllers"

	"github.com/alex-bezverkhniy/gonotes/routers"
)

// App - Application
type App struct {
	Port           string
	DataFileName   string
	Router         *routers.Router
	NoteController *controllers.NoteController
}

// NewApp - create new instance
func NewApp(port, dataFileName string) *App {
	a := &App{Port: port, DataFileName: dataFileName}
	noteRepository := repositories.NewNoteRepository(a.DataFileName)
	a.NoteController = controllers.NewNoteController(noteRepository)
	a.Router = routers.NewRouter(http.NewServeMux(), a.NoteController)

	return a
}

// Serve - start http server
func (a *App) Serve() {
	log.Println("Go Notes Rest API v1.0")
	if err := http.ListenAndServe(a.Port, a.Router.Mux); err != nil {
		log.Fatal(err)
	}
}

// Shutdown - stops http server
func (a *App) Shutdown(code int) {
	os.Exit(code)
}
