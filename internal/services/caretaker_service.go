// Package services provides the business logic for the caretaker operations.
package services

import (
	"context"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/repository"
)

type CaretakerService interface {
	CreateCaretaker(ctx context.Context, caretaker *models.Caretaker) error
	GetCaretaker(ctx context.Context, id string) (*models.Caretaker, error)
	UpdateCaretaker(ctx context.Context, caretaker *models.Caretaker) error
	DeleteCaretaker(ctx context.Context, id string) error
	ListCaretakers(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Caretaker, error)
}

type caretakerService struct {
	repo repository.CaretakerRepository
}

func NewCaretakerService(repo repository.CaretakerRepository) CaretakerService {
	return &caretakerService{
		repo: repo,
	}
}

func (s *caretakerService) CreateCaretaker(ctx context.Context, caretaker *models.Caretaker) error {
	return s.repo.Create(ctx, caretaker)
}

func (s *caretakerService) GetCaretaker(ctx context.Context, id string) (*models.Caretaker, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *caretakerService) UpdateCaretaker(ctx context.Context, caretaker *models.Caretaker) error {
	return s.repo.Update(ctx, caretaker)
}

func (s *caretakerService) DeleteCaretaker(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *caretakerService) ListCaretakers(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Caretaker, error) {
	return s.repo.List(ctx, filter, page, pageSize)
}
