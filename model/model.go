package model

import (
	"fmt"
	"time"
)

// Note - Structure for notes
type Note struct {
	ID        int       `json:id`
	Title     string    `json:title`
	Content   string    `json:content`
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
}

func (n *Note) String() string {
	return fmt.Sprintf("{ID: %d, Title: %s, Content: %s, CreatedAt: %v, UpdatedAt: %v}",
		n.ID,
		n.Title,
		n.Content,
		n.CreatedAt,
		n.UpdatedAt,
	)
}
