// Package repository provides data access layer implementations.
package repository

import (
	"context"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type NeedRepository interface {
	Create(ctx context.Context, need *models.Need) error
	GetByID(ctx context.Context, id string) (*models.Need, error)
}

type needRepository struct {
	db *sqlx.DB
}

func NewNeedRepository(db *sqlx.DB) NeedRepository {
	return &needRepository{db: db}
}

func (r *needRepository) Create(ctx context.Context, need *models.Need) error {
	const query = `
		INSERT INTO needs (type, description)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		need.Type,
		need.Description,
	).Scan(&need.ID, &need.CreatedAt, &need.UpdatedAt)
}

func (r *needRepository) GetByID(ctx context.Context, id string) (*models.Need, error) {
	const query = `SELECT id, type, description, created_at, updated_at FROM needs WHERE id = $1`

	var need models.Need
	err := r.db.GetContext(ctx, &need, query, id)
	if err != nil {
		return nil, fmt.Errorf("needRepository.GetByID: %w", err)
	}
	return &need, nil
}
