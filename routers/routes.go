package routers

import (
	"net/http"

	"github.com/alex-bezverkhniy/gonotes/controllers"
)

// Router - Router abstraction
type Router struct {
	Mux            *http.ServeMux
	NoteController *controllers.NoteController
}

// NewRouter - creates new instance
func NewRouter(mux *http.ServeMux, nc *controllers.NoteController) *Router {
	r := &Router{Mux: mux, NoteController: nc}

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/notes/", nc.Dispatch)
	mux.HandleFunc("/notes", nc.Dispatch)
	mux.HandleFunc("/flush/notes", nc.Flush)

	return r
}
