package model

import (
	"time"
)

// Access struct
type Access struct {
	IP        string    `json:"ip" bson:"ip"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
