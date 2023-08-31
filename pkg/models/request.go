package models

type SegmentRequest struct {
	Slug string
}

type ManageUserToSegmentsRequest struct {
	SlugsToAdd    []string `json:"slugs-to-add" validate:"empty=false"`
	SlugsToRemove []string `json:"slugs-to-remove" validate:"empty=false"`
	UserId        uint     `json:"user-id" validate:"empty=false"`
}

type GetUserSegmentsRequest struct {
	UserId uint `json:"user-id" validate:"empty=false"`
}
