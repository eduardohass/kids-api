// internal/models/group.go
package models

import (
	"time"
)

type Group struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	AgeRange    string    `json:"age_range" db:"age_range"`
	Capacity    int       `json:"capacity" db:"capacity"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
