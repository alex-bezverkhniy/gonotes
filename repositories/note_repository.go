package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/alex-bezverkhniy/gonotes/model"
)

// NoteRepository - Repository for notes
type NoteRepository struct {
	DataFileName string
	Notes        []model.Note
}

// NewNoteRepository - Creates new instance of repository
func NewNoteRepository(dataFileName string) *NoteRepository {
	return &NoteRepository{
		DataFileName: dataFileName,
		Notes:        loadNotes(dataFileName),
	}

}

// FindAll - Returns all notes
func (nr *NoteRepository) FindAll() []model.Note {
	if nr.Notes == nil {
		nr.Notes = loadNotes(nr.DataFileName)
	}
	return nr.Notes
}

// FindByID - Returns note by ID
func (nr *NoteRepository) FindByID(ID string) (model.Note, error) {
	for _, note := range nr.Notes {
		if note.ID == ID {
			return note, nil
		}
	}
	return model.Note{}, errors.New("note NOT found")
}

// Create - Creates new note
func (nr *NoteRepository) Create(note model.Note) (model.Note, error) {
	nr.Notes = append(nr.Notes, note)
	return note, nil
}

// Update - Creates new note
func (nr *NoteRepository) Update(ID string, note model.Note) (model.Note, error) {
	index := -1
	for i, note := range nr.Notes {
		if note.ID == ID {
			index = i
		}
	}

	if index == -1 {
		return model.Note{}, errors.New("note NOT found")
	}

	nr.Notes[index] = note

	return note, nil
}

// Delete - Removes note by ID
func (nr *NoteRepository) Delete(ID string) error {
	log.Println("NoteRepository.Delete")
	index := -1
	for i, note := range nr.Notes {
		if note.ID == ID {
			index = i
			nr.Notes = append(nr.Notes[:i], nr.Notes[i+1:]...)
		}
	}
	if index == -1 {
		return errors.New("note NOT found")
	} else {
		return nil
	}

}

func loadNotes(dataFileName string) []model.Note {
	file, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		log.Panicln("Error of reading json data: ", err)
		return nil
	}

	data := model.NotesList{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		log.Panicln("Error of Unmarshaling json data: ", err)
		return nil
	}

	return data
}
