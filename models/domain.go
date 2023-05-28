package models

import (
	"github.com/google/uuid"
)

type Domain struct {
	Id      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
	Label   string
}
