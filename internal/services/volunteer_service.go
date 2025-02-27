package services

import (
	"context"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/repository"
)

type VolunteerService interface {
	CreateVolunteer(ctx context.Context, volunteer *models.Volunteer) error
	GetVolunteer(ctx context.Context, id string) (*models.Volunteer, error)
	UpdateVolunteer(ctx context.Context, volunteer *models.Volunteer) error
	DeleteVolunteer(ctx context.Context, id string) error
	ListVolunteers(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Volunteer, error)
}

type volunteerService struct {
	repo repository.VolunteerRepository
}

func NewVolunteerService(repo repository.VolunteerRepository) VolunteerService {
	return &volunteerService{
		repo: repo,
	}
}

func (s *volunteerService) CreateVolunteer(ctx context.Context, volunteer *models.Volunteer) error {
	return s.repo.Create(ctx, volunteer)
}

func (s *volunteerService) GetVolunteer(ctx context.Context, id string) (*models.Volunteer, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *volunteerService) UpdateVolunteer(ctx context.Context, volunteer *models.Volunteer) error {
	return s.repo.Update(ctx, volunteer)
}

func (s *volunteerService) DeleteVolunteer(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *volunteerService) ListVolunteers(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Volunteer, error) {
	return s.repo.List(ctx, filter, page, pageSize)
}
