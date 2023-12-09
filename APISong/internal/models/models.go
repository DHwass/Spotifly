package models

import (
	"github.com/gofrs/uuid"
)

// On garde cet exemple pour le moment
type Collection struct {
	Id      *uuid.UUID `json:"id"`
	Content string     `json:"content"`
}

// Song model
type Song struct {
	Id       *uuid.UUID `json:"id"`
	Title    string     `json:"title"`
	Artist   string     `json:"artist"`
	Duration int        `json:"duration"`
}
