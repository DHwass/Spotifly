package models

import (
	"github.com/gofrs/uuid"
)

type Users struct {
	Id    *uuid.UUID `json:"id"`
	Name  string     `json:"name"`
	Email string     `json:"email"`
}
