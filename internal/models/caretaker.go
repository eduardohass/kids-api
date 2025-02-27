// internal/models/caretaker.go
// Package models defines the data structures used throughout the application.
package models

import (
	"time"
)

// Caretaker represents a person responsible for one or more children.
type Caretaker struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Address   string    `json:"address" db:"address"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ChildCaretakerRelation represents the relationship between a child and their caretaker.
type ChildCaretakerRelation struct {
	ID           string    `json:"id" db:"id"`
	ChildID      string    `json:"child_id" db:"crianca_id"`
	CaretakerID  string    `json:"caretaker_id" db:"responsavel_id"`
	RelationType string    `json:"relation_type" db:"tipo_relacao"`
	CanPickup    bool      `json:"can_pickup" db:"pode_retirar"`
	CreatedAt    time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt    time.Time `json:"updated_at" db:"atualizado_em"`
}
