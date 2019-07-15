package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/alex-bezverkhniy/gonotes/model"
	"github.com/alex-bezverkhniy/gonotes/repositories"
)

// NoteController - HTTP controller for notes
type NoteController struct {
	NoteRepository *repositories.NoteRepository
}

// NewNoteController - Creates new instance of controller
func NewNoteController(noteRepository *repositories.NoteRepository) *NoteController {
	return &NoteController{
		NoteRepository: noteRepository,
	}
}

func (nc *NoteController) Dispatch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nc.Create(w, r)
		return
	}

	if r.Method == "PUT" {
		nc.Update(w, r)
		return
	}

	if r.Method == "DELETE" {
		log.Println(r.URL.Path)
		nc.Delete(w, r)
		return
	}

	if r.Method == "GET" {
		if strings.HasSuffix(r.URL.Path, "/notes") {
			nc.FindAll(w, r)
		} else {
			nc.FindByID(w, r)
		}

		return
	}
}

// FindAll - Returns all notes
func (nc *NoteController) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Println("FindAll")
	notes := nc.NoteRepository.FindAll()
	json.NewEncoder(w).Encode(notes)
}

// FindByID - Returns note by ID
func (nc *NoteController) FindByID(w http.ResponseWriter, r *http.Request) {
	log.Println("FindByID")
	ID := path.Base(r.URL.Path)
	if ID == " " {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	note, err := nc.NoteRepository.FindByID(ID)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// Create - Creates new note
func (nc *NoteController) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("Create")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Format is not correct", http.StatusBadRequest)
		return
	}

	var note model.Note

	json.Unmarshal(reqBody, &note)
	note.CreatedAt = time.Now()

	note, err = nc.NoteRepository.Create(note)

	if err != nil {
		http.Error(w, "Error during saving note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// Update - Creates new note
func (nc *NoteController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update")
	ID := path.Base(r.URL.Path)
	if ID == " " {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Format is not correct", http.StatusBadRequest)
		return
	}

	var note model.Note

	json.Unmarshal(reqBody, &note)
	note.CreatedAt = time.Now()

	note, err = nc.NoteRepository.Update(ID, note)

	if err != nil {
		log.Panicln("Error during updating note: ", err)
		http.Error(w, "Error during updating note", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(note)
}

// Delete - Removes note by ID
func (nc *NoteController) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete")
	ID := path.Base(r.URL.Path)
	if ID == " " {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	log.Println("ID: ", ID)

	err := nc.NoteRepository.Delete(ID)
	if err != nil {
		log.Panicln("Error deleting note: ", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
