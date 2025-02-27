// Package models provides the definitions for the models used in the application.
package models

import "time"

type Volunteer struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	Skills       string    `json:"skills" db:"skills"`
	Availability string    `json:"availability" db:"availability"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
