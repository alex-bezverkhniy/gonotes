package model

import "time"

// Note - datatype tor notes
type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Desc      string    `json:"desc"`
	Content   string    `json:"content"`
}

// NotesList - list of notes
type NotesList []Note
