package routers

import (
	"net/http"

	"github.com/alex-bezverkhniy/gonotes/controllers"
)

// CreateRoutes - Creates new router
func CreateRoutes(mux *http.ServeMux, nc *controllers.NoteController) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	mux.HandleFunc("/notes/", nc.Dispatch)
	mux.HandleFunc("/notes", nc.Dispatch)
}
