package repository

import (
	"backend-trainee-assignment-2023/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Segment interface {
	Create(slug string) (uint, error)
	Delete(slug string) error
}

type User interface {
	ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error)
	GetUserSegments(userId uint) ([]string, error)
}

type Repository struct {
	Segment
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Segment: NewSegmentRepository(db),
		User:    NewUsersSegmentsRepository(db),
	}
}
