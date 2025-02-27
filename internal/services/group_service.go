package services

import (
	"context"

	"github.com/eduardohass/kids-api/internal/models"
	"github.com/eduardohass/kids-api/internal/repository"
)

type GroupService interface {
	CreateGroup(ctx context.Context, group *models.Group) error
	GetGroup(ctx context.Context, id string) (*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, id string) error
	ListGroups(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Group, error)
}

type groupService struct {
	repo repository.GroupRepository
}

func NewGroupService(repo repository.GroupRepository) GroupService {
	return &groupService{
		repo: repo,
	}
}

func (s *groupService) CreateGroup(ctx context.Context, group *models.Group) error {
	return s.repo.Create(ctx, group)
}

func (s *groupService) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *groupService) UpdateGroup(ctx context.Context, group *models.Group) error {
	return s.repo.Update(ctx, group)
}

func (s *groupService) DeleteGroup(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *groupService) ListGroups(ctx context.Context, filter map[string]interface{}, page, pageSize int) ([]*models.Group, error) {
	return s.repo.List(ctx, filter, page, pageSize)
}
