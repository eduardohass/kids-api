// Package repository provides data access layer implementations.
package repository

import (
	"context"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type AllergyRepository interface {
	Create(ctx context.Context, allergy *models.Allergy) error
	GetByID(ctx context.Context, id string) (*models.Allergy, error)
}

type allergyRepository struct {
	db *sqlx.DB
}

func NewAllergyRepository(db *sqlx.DB) AllergyRepository {
	return &allergyRepository{db: db}
}

func (r *allergyRepository) Create(ctx context.Context, allergy *models.Allergy) error {
	// TODO: implemente a criação do registro de allergy conforme necessário
	return nil
}

func (r *allergyRepository) GetByID(ctx context.Context, id string) (*models.Allergy, error) {
	const query = `SELECT id, type, description, severity, created_at, updated_at FROM allergies WHERE id = $1`

	var allergy models.Allergy
	err := r.db.GetContext(ctx, &allergy, query, id)
	if err != nil {
		return nil, fmt.Errorf("allergyRepository.GetByID: %w", err)
	}
	return &allergy, nil
}
