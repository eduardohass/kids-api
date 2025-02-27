package repository

import (
	"context"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type CaretakerRepository interface {
	Create(ctx context.Context, caretaker *models.Caretaker) error
	GetByID(ctx context.Context, id string) (*models.Caretaker, error)
	Update(ctx context.Context, caretaker *models.Caretaker) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Caretaker, error)
}

type caretakerRepository struct {
	db *sqlx.DB
}

func NewCaretakerRepository(db *sqlx.DB) CaretakerRepository {
	return &caretakerRepository{db: db}
}

func (r *caretakerRepository) Create(ctx context.Context, caretaker *models.Caretaker) error {
	const query = `
		INSERT INTO caretakers (
			name,
			email,
			phone,
			address
		) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		caretaker.Name,
		caretaker.Email,
		caretaker.Phone,
		caretaker.Address,
	).Scan(&caretaker.ID, &caretaker.CreatedAt, &caretaker.UpdatedAt)
}

func (r *caretakerRepository) GetByID(ctx context.Context, id string) (*models.Caretaker, error) {
	const query = `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM caretakers
		WHERE id = $1
	`

	var caretaker models.Caretaker
	if err := r.db.GetContext(ctx, &caretaker, query, id); err != nil {
		return nil, fmt.Errorf("caretakerRepository.GetByID: %w", err)
	}
	return &caretaker, nil
}

func (r *caretakerRepository) Update(ctx context.Context, caretaker *models.Caretaker) error {
	const query = `
		UPDATE caretakers SET
			name = $1,
			email = $2,
			phone = $3,
			address = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		caretaker.Name,
		caretaker.Email,
		caretaker.Phone,
		caretaker.Address,
		caretaker.ID,
	)
	if err != nil {
		return fmt.Errorf("caretakerRepository.Update: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("caretakerRepository.Update: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *caretakerRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM caretakers WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("caretakerRepository.Delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("caretakerRepository.Delete: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *caretakerRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Caretaker, error) {
	const query = `
		SELECT id, name, email, phone, address, created_at, updated_at
		FROM caretakers
		LIMIT $1 OFFSET $2
	`

	var caretakers []*models.Caretaker
	if err := r.db.SelectContext(ctx, &caretakers, query, limit, offset); err != nil {
		return nil, fmt.Errorf("caretakerRepository.List: %w", err)
	}

	return caretakers, nil
}
