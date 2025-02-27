// internal/services/child_service.go
package services

import (
	"context"
	"errors"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/repository"
)

type ChildService interface {
	CreateChild(ctx context.Context, child *models.Child) error
	GetChild(ctx context.Context, id string) (*models.Child, error)
	UpdateChild(ctx context.Context, child *models.Child) error
	DeleteChild(ctx context.Context, id string) error
	ListChildren(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Child, error)
}

type childService struct {
	childRepo   repository.ChildRepository
	needRepo    repository.NeedRepository
	allergyRepo repository.AllergyRepository
}

func NewChildService(
	childRepo repository.ChildRepository,
	needRepo repository.NeedRepository,
	allergyRepo repository.AllergyRepository,
) ChildService {
	return &childService{
		childRepo:   childRepo,
		needRepo:    needRepo,
		allergyRepo: allergyRepo,
	}
}

func (s *childService) CreateChild(ctx context.Context, child *models.Child) error {
	// Validações
	if child.Name == "" {
		return errors.New("child name is required")
	}

	if child.BirthDate.IsZero() {
		return errors.New("birth date is required")
	}

	if child.Gender == "" {
		return errors.New("gender is required")
	}

	// Verificar se necessidades existem
	for i, need := range child.Needs {
		if need.ID == "" {
			// Criar necessidade se não existir
			err := s.needRepo.Create(ctx, &need)
			if err != nil {
				return err
			}
			child.Needs[i] = need
		}
	}

	// Verificar se alergias existem
	for i, allergy := range child.Allergies {
		if allergy.ID == "" {
			// Criar alergia se não existir
			err := s.allergyRepo.Create(ctx, &allergy)
			if err != nil {
				return err
			}
			child.Allergies[i] = allergy
		}
	}

	return s.childRepo.Create(ctx, child)
}

func (s *childService) DeleteChild(ctx context.Context, id string) error {
	return s.childRepo.Delete(ctx, id)
}

func (s *childService) GetChild(ctx context.Context, id string) (*models.Child, error) {
	return s.childRepo.GetByID(ctx, id)
}

func (s *childService) UpdateChild(ctx context.Context, child *models.Child) error {
	return s.childRepo.Update(ctx, child)
}

func (s *childService) ListChildren(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Child, error) {
	return s.childRepo.List(ctx, filter, page, pageSize)
}
