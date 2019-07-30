package main

import (
	"github.com/alex-bezverkhniy/gonotes/model"
)

// NotesStore - basic operations with note
type NotesStore interface {
	Save(note model.Note)
	Get(id int)
}
