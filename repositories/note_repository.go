package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/alex-bezverkhniy/gonotes/model"
)

// NoteRepository - Repository for notes
type NoteRepository struct {
	DataFileName string
	Notes        map[string]model.Note
}

// NewNoteRepository - Creates new instance of repository
func NewNoteRepository(dataFileName string) *NoteRepository {
	return &NoteRepository{
		DataFileName: dataFileName,
		Notes:        loadNotes(dataFileName),
	}

}

// FindAll - Returns all notes
func (nr *NoteRepository) FindAll() map[string]model.Note {
	if nr.Notes == nil {
		nr.Notes = loadNotes(nr.DataFileName)
	}
	return nr.Notes
}

// FindByID - Returns note by ID
func (nr *NoteRepository) FindByID(ID string) (model.Note, error) {
	if ID == "" {
		return model.Note{}, errors.New("ID should NOT be empty")
	}

	note := nr.Notes[ID]
	if (note == model.Note{}) {
		return model.Note{}, errors.New("note NOT found")
	}

	return note, nil
}

// Create - Creates new note
func (nr *NoteRepository) Create(note model.Note) (model.Note, error) {
	if note.ID == "" {
		return model.Note{}, errors.New("ID should NOT be empty")
	}

	n, _ := nr.FindByID(note.ID)
	if (n != model.Note{}) {
		return model.Note{}, errors.New("note already exists")
	}

	nr.Notes[note.ID] = note
	go storeNotes(nr.DataFileName, nr.Notes)
	return note, nil
}

// Update - Creates new note
func (nr *NoteRepository) Update(ID string, note model.Note) (model.Note, error) {
	if ID == "" {
		return model.Note{}, errors.New("ID should NOT be empty")
	}

	if (note == model.Note{}) {
		return model.Note{}, errors.New("note should NOT be empty")
	}

	if (nr.Notes[ID] == model.Note{}) {
		return model.Note{}, errors.New("note NOT found")
	}

	nr.Notes[ID] = note
	go storeNotes(nr.DataFileName, nr.Notes)
	return note, nil
}

// Delete - Removes note by ID
func (nr *NoteRepository) Delete(ID string) error {
	if ID == "" {
		return errors.New("ID should NOT be empty")
	}

	if (nr.Notes[ID] == model.Note{}) {
		return errors.New("note NOT found")
	}
	delete(nr.Notes, ID)
	go storeNotes(nr.DataFileName, nr.Notes)
	return nil
}

// Flush - load data from the file
func (nr *NoteRepository) Flush() {
	nr.Notes = loadNotes(nr.DataFileName)
}

func loadNotes(dataFileName string) map[string]model.Note {
	file, err := ioutil.ReadFile(dataFileName)
	if err != nil {
		log.Panicln("Error of reading json data: ", err)
		return nil
	}

	data := map[string]model.Note{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		log.Panicln("Error of Unmarshaling json data: ", err)
		return nil
	}

	return data
}

func storeNotes(dataFileName string, data map[string]model.Note) error {
	if data == nil || len(data) <= 0 {
		return errors.New("data for storing should not be empty")
	}

	if _, err := os.Stat(dataFileName); os.IsNotExist(err) {
		// Create file
		_, err = os.Create(dataFileName)

		if err != nil {
			return err
		}
	}
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dataFileName, bytesData, 0644) //file.Write(bytesData)
	if err != nil {
		return err
	}

	return nil
}
