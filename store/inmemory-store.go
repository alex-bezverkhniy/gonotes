package store

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"github.com/alex-bezverkhniy/gonotes/model"
)

var errEmptyBuffer = errors.New("cannot read from empty buffer")

// InMemoryNoteStore - structure for storing data in memory
type InMemoryNoteStore struct {
	store map[int]model.Note
}

// DataLoader - operatons for loading data into memmory
type DataLoader interface {
	Load() (map[int]model.Note, error)
}

// FileDataLoader - file loader
type FileDataLoader struct {
	reader *io.Reader
}

// NewFileDataLoader - creates new instance of FileDataLoader
func NewFileDataLoader(reader io.Reader) *FileDataLoader {
	return &FileDataLoader{&reader}
}

// Load - loads initial data from file
func (f *FileDataLoader) Load() (map[int]model.Note, error) {
	buf, _ := ioutil.ReadAll(*f.reader)
	data := map[int]model.Note{}
	err := json.Unmarshal(buf, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInMemoryNoteStore - creates new instance of InMemoryNoteStore
func NewInMemoryNoteStore(loader DataLoader) *InMemoryNoteStore {
	instance := &InMemoryNoteStore{store: map[int]model.Note{}}
	instance.store, _ = loader.Load()
	return instance
}

// Get returns note by ID
func (i *InMemoryNoteStore) Get(id int) model.Note {
	return i.store[id]
}
