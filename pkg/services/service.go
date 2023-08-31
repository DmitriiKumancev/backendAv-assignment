package services

import (
	"backend-trainee-assignment-2023/pkg/models"
	"backend-trainee-assignment-2023/pkg/repository"
)

type Segment interface {
	Create(slug string) (uint, error)
	Delete(slug string) error
}

type User interface {
	ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error)
	GetUserSegments(userId uint) ([]string, error)
}

type Service struct {
	Segment
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Segment: NewSegmentService(repo.Segment),
		User:    NewUsersSegmentsService(repo.User),
	}
}
