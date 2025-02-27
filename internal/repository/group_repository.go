// Package repository provides data access layer implementations.
package repository

import (
	"context"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type GroupRepository interface {
	Create(ctx context.Context, group *models.Group) error
	GetByID(ctx context.Context, id string) (*models.Group, error)
	Update(ctx context.Context, group *models.Group) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Group, error)
}

type groupRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (r *groupRepository) Create(ctx context.Context, group *models.Group) error {
	const query = `
		INSERT INTO groups (
			name,
			description,
			age_range,
			capacity
		) VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRowxContext(
		ctx,
		query,
		group.Name,
		group.Description,
		group.AgeRange,
		group.Capacity,
	).Scan(&group.ID, &group.CreatedAt, &group.UpdatedAt)
}

func (r *groupRepository) GetByID(ctx context.Context, id string) (*models.Group, error) {
	const query = `
		SELECT id, name, description, age_range, capacity, created_at, updated_at
		FROM groups
		WHERE id = $1
	`

	var group models.Group
	if err := r.db.GetContext(ctx, &group, query, id); err != nil {
		return nil, fmt.Errorf("groupRepository.GetByID: %w", err)
	}
	return &group, nil
}

func (r *groupRepository) Update(ctx context.Context, group *models.Group) error {
	const query = `
		UPDATE groups SET
			name = $1,
			description = $2,
			age_range = $3,
			capacity = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING updated_at
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		group.Name,
		group.Description,
		group.AgeRange,
		group.Capacity,
		group.ID,
	)
	if err != nil {
		return fmt.Errorf("groupRepository.Update: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("groupRepository.Update: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *groupRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM groups WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("groupRepository.Delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("groupRepository.Delete: %w", err)
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (r *groupRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Group, error) {
	const query = `
		SELECT id, name, description, age_range, capacity, created_at, updated_at
		FROM groups
		LIMIT $1 OFFSET $2
	`

	var groups []*models.Group
	if err := r.db.SelectContext(ctx, &groups, query, limit, offset); err != nil {
		return nil, fmt.Errorf("groupRepository.List: %w", err)
	}

	return groups, nil
}
