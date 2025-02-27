package repository

import (
	"context"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type VolunteerRepository interface {
	Create(ctx context.Context, volunteer *models.Volunteer) error
	GetByID(ctx context.Context, id string) (*models.Volunteer, error)
	Update(ctx context.Context, volunteer *models.Volunteer) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Volunteer, error)
}

type volunteerRepository struct {
	db *sqlx.DB
}

func NewVolunteerRepository(db *sqlx.DB) VolunteerRepository {
	return &volunteerRepository{db: db}
}

func (r *volunteerRepository) Create(ctx context.Context, volunteer *models.Volunteer) error {
	const query = `
		INSERT INTO volunteers (
			name,
			email,
			phone,
			skills,
			availability
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		volunteer.Name,
		volunteer.Email,
		volunteer.Phone,
		volunteer.Skills,
		volunteer.Availability,
	).Scan(&volunteer.ID, &volunteer.CreatedAt, &volunteer.UpdatedAt)
}

func (r *volunteerRepository) GetByID(ctx context.Context, id string) (*models.Volunteer, error) {
	const query = `
		SELECT id, name, email, phone, skills, availability, created_at, updated_at
		FROM volunteers
		WHERE id = $1
	`

	var volunteer models.Volunteer
	if err := r.db.GetContext(ctx, &volunteer, query, id); err != nil {
		return nil, fmt.Errorf("volunteerRepository.GetByID: %w", err)
	}
	return &volunteer, nil
}

func (r *volunteerRepository) Update(ctx context.Context, volunteer *models.Volunteer) error {
	const query = `
		UPDATE volunteers SET
			name = $1,
			email = $2,
			phone = $3,
			skills = $4,
			availability = $5,
			updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		volunteer.Name,
		volunteer.Email,
		volunteer.Phone,
		volunteer.Skills,
		volunteer.Availability,
		volunteer.ID,
	)
	if err != nil {
		return fmt.Errorf("volunteerRepository.Update: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("volunteerRepository.Update: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *volunteerRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM volunteers WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("volunteerRepository.Delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("volunteerRepository.Delete: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *volunteerRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Volunteer, error) {
	const query = `
		SELECT id, name, email, phone, skills, availability, created_at, updated_at
		FROM volunteers
		LIMIT $1 OFFSET $2
	`

	var volunteers []*models.Volunteer
	if err := r.db.SelectContext(ctx, &volunteers, query, limit, offset); err != nil {
		return nil, fmt.Errorf("volunteerRepository.List: %w", err)
	}

	return volunteers, nil
}
