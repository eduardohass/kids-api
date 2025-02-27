// Package repository provides data access layer implementations.
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/jmoiron/sqlx"
)

// ErrNotFound is returned when a requested entity is not found in the database.
var ErrNotFound = errors.New("child not found")

type ChildRepository interface {
	Create(ctx context.Context, child *models.Child) error
	GetByID(ctx context.Context, id string) (*models.Child, error)
	Update(ctx context.Context, child *models.Child) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Child, error)
	AssociateNeed(ctx context.Context, childID, needID string) error
	AssociateAllergy(ctx context.Context, childID, allergyID string) error
}

type childRepository struct {
	db *sqlx.DB
}

func NewChildRepository(db *sqlx.DB) ChildRepository {
	return &childRepository{db: db}
}

func (r *childRepository) Create(ctx context.Context, child *models.Child) error {
	const query = `
		INSERT INTO children (
			name, 
			birth_date, 
			gender, 
			photo_url, 
			group_id
		) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowxContext(
		ctx,
		query,
		child.Name,
		child.BirthDate,
		child.Gender,
		child.PhotoURL,
		child.GroupID,
	).Scan(&child.ID, &child.CreatedAt, &child.UpdatedAt)

	if err != nil {
		return fmt.Errorf("childRepository.Create: %w", err)
	}

	if err := r.associateRelatedEntities(ctx, child); err != nil {
		return fmt.Errorf("childRepository.Create: %w", err)
	}

	return nil
}

func (r *childRepository) GetByID(ctx context.Context, id string) (*models.Child, error) {
	const baseQuery = `
		SELECT 
			id, 
			name, 
			birth_date, 
			gender, 
			photo_url, 
			group_id, 
			created_at, 
			updated_at
		FROM children
		WHERE id = $1
	`

	var child models.Child
	if err := r.db.GetContext(ctx, &child, baseQuery, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("childRepository.GetByID: %w", err)
	}

	if err := r.loadAssociations(ctx, &child); err != nil {
		return nil, fmt.Errorf("childRepository.GetByID: %w", err)
	}

	return &child, nil
}

func (r *childRepository) Update(ctx context.Context, child *models.Child) error {
	const query = `
		UPDATE children SET
			name = $1,
			birth_date = $2,
			gender = $3,
			photo_url = $4,
			group_id = $5,
			updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		child.Name,
		child.BirthDate,
		child.Gender,
		child.PhotoURL,
		child.GroupID,
		child.ID,
	); err != nil {
		return fmt.Errorf("childRepository.Update: %w", err)
	}

	if err := r.syncAssociations(ctx, child); err != nil {
		return fmt.Errorf("childRepository.Update: %w", err)
	}

	return nil
}

func (r *childRepository) Delete(ctx context.Context, id string) error {
	const query = `DELETE FROM children WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("childRepository.Delete: %w", err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return ErrNotFound
	}

	return nil
}

// Adição da implementação do método List para satisfazer a interface
func (r *childRepository) List(ctx context.Context, filter map[string]interface{}, limit, offset int) ([]*models.Child, error) {
	const query = `
		SELECT 
			id, 
			name, 
			birth_date, 
			gender, 
			photo_url, 
			group_id, 
			created_at, 
			updated_at
		FROM children
		LIMIT $1 OFFSET $2
	`
	var children []*models.Child
	if err := r.db.SelectContext(ctx, &children, query, limit, offset); err != nil {
		return nil, fmt.Errorf("childRepository.List: %w", err)
	}
	return children, nil
}

// Implementações auxiliares
func (r *childRepository) associateRelatedEntities(ctx context.Context, child *models.Child) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := r.associateNeeds(ctx, tx, child); err != nil {
		return err
	}

	if err := r.associateAllergies(ctx, tx, child); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *childRepository) associateNeeds(ctx context.Context, tx *sqlx.Tx, child *models.Child) error {
	for _, need := range child.Needs {
		if err := r.associateNeed(ctx, tx, child.ID, need.ID); err != nil {
			return err
		}
	}
	return nil
}

func (r *childRepository) associateAllergies(ctx context.Context, tx *sqlx.Tx, child *models.Child) error {
	for _, allergy := range child.Allergies {
		if err := r.associateAllergy(ctx, tx, child.ID, allergy.ID); err != nil {
			return err
		}
	}
	return nil
}

func (r *childRepository) loadAssociations(ctx context.Context, child *models.Child) error {
	if err := r.loadNeeds(ctx, child); err != nil {
		return err
	}
	return r.loadAllergies(ctx, child)
}

func (r *childRepository) loadNeeds(ctx context.Context, child *models.Child) error {
	const query = `
		SELECT 
			n.id,
			n.type,
			n.description,
			n.created_at,
			n.updated_at
		FROM needs n
		INNER JOIN child_needs cn ON n.id = cn.need_id
		WHERE cn.child_id = $1
	`
	return r.db.SelectContext(ctx, &child.Needs, query, child.ID)
}

func (r *childRepository) loadAllergies(ctx context.Context, child *models.Child) error {
	const query = `
		SELECT 
			a.id,
			a.type,
			a.description,
			a.severity,
			a.created_at,
			a.updated_at
		FROM allergies a
		INNER JOIN child_allergies ca ON a.id = ca.allergy_id
		WHERE ca.child_id = $1
	`
	return r.db.SelectContext(ctx, &child.Allergies, query, child.ID)
}

func (r *childRepository) syncAssociations(ctx context.Context, child *models.Child) error {
	// Implementação de sincronização de associações
	// (Atualização diferencial das necessidades e alergias)
	return nil
}

// Métodos de associação mantendo a transação
func (r *childRepository) AssociateNeed(ctx context.Context, childID, needID string) error {
	return r.associateNeed(ctx, r.db, childID, needID)
}

func (r *childRepository) associateNeed(ctx context.Context, exec sqlx.ExecerContext, childID, needID string) error {
	const query = `
		INSERT INTO child_needs (child_id, need_id)
		VALUES ($1, $2)
		ON CONFLICT (child_id, need_id) DO NOTHING
	`
	_, err := exec.ExecContext(ctx, query, childID, needID)
	return err
}

func (r *childRepository) AssociateAllergy(ctx context.Context, childID, allergyID string) error {
	return r.associateAllergy(ctx, r.db, childID, allergyID)
}

func (r *childRepository) associateAllergy(ctx context.Context, exec sqlx.ExecerContext, childID, allergyID string) error {
	const query = `
		INSERT INTO child_allergies (child_id, allergy_id)
		VALUES ($1, $2)
		ON CONFLICT (child_id, allergy_id) DO NOTHING
	`
	_, err := exec.ExecContext(ctx, query, childID, allergyID)
	return err
}
