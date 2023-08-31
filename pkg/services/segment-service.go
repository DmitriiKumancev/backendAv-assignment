package services

import "backend-trainee-assignment-2023/pkg/repository"

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (service *SegmentService) Create(slug string) (uint, error) {
	return service.repo.Create(slug)
}

func (service *SegmentService) Delete(slug string) error {
	return service.repo.Delete(slug)
}
