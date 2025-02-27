// internal/models/child.go
package models

import (
	"time"
)

type Need struct {
	ID          string    `json:"id" db:"id"`
	Type        string    `json:"type" db:"tipo"`
	Description string    `json:"description" db:"descricao"`
	CreatedAt   time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt   time.Time `json:"updated_at" db:"atualizado_em"`
}

type Allergy struct {
	ID          string    `json:"id" db:"id"`
	Type        string    `json:"type" db:"tipo"`
	Description string    `json:"description" db:"descricao"`
	Severity    string    `json:"severity" db:"gravidade"`
	CreatedAt   time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt   time.Time `json:"updated_at" db:"atualizado_em"`
}

type Child struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"nome"`
	BirthDate time.Time `json:"birth_date" db:"data_nascimento"`
	Gender    string    `json:"gender" db:"sexo"`
	PhotoURL  string    `json:"photo_url" db:"foto_url"`
	Needs     []Need    `json:"needs"`
	Allergies []Allergy `json:"allergies"`
	GroupID   string    `json:"group_id" db:"grupo_id"`
	CreatedAt time.Time `json:"created_at" db:"criado_em"`
	UpdatedAt time.Time `json:"updated_at" db:"atualizado_em"`
}
