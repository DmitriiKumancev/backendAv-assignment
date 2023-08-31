package repository

import (
	"backend-trainee-assignment-2023/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type UsersSegmentsRepository struct {
	db *sqlx.DB
}

type operation string

const (
	add    operation = "add"
	delete operation = "delete"
)

func NewUsersSegmentsRepository(db *sqlx.DB) *UsersSegmentsRepository {
	return &UsersSegmentsRepository{db: db}
}

func (repo *UsersSegmentsRepository) ManageUserToSegments(slugsToAdd []string, slugsToRemove []string, userId uint) (*models.ManageUserToSegmentsResponse, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(fmt.Sprintf("SET TRANSACTION ISOLATION LEVEL %s", "REPEATABLE READ"))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	segmentsToAdd, err := repo.filterSegments(slugsToAdd)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	segmentsToRemove, err := repo.filterSegments(slugsToRemove)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(segmentsToAdd) != 0 {
		if err = repo.insertSegmentsIntoUser(segmentsToAdd, userId); err != nil {
			tx.Rollback()
			return nil, err
		}
		if err = repo.saveInHistory(segmentsToAdd, userId, add); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if len(segmentsToRemove) != 0 {
		if err = repo.deleteSegmentsFromUser(segmentsToRemove, userId); err != nil {
			tx.Rollback()
			return nil, err
		}
		if err = repo.saveInHistory(segmentsToRemove, userId, delete); err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	_, slugsHaveBeenAdded := decomposeSegments(segmentsToAdd)
	_, slugsHaveBeenRemoved := decomposeSegments(segmentsToRemove)
	return &models.ManageUserToSegmentsResponse{
		SlugsHaveBeenAdded:   slugsHaveBeenAdded,
		SlugsHaveBeenRemoved: slugsHaveBeenRemoved,
		UserId:               userId,
	}, tx.Commit()
}

func (repo *UsersSegmentsRepository) GetUserSegments(userId uint) ([]string, error) {
	var slugs []string
	query := fmt.Sprintf(
		`SELECT s.slug FROM %s AS us 
                JOIN %s AS s ON s.id=us.segment_id 
            	WHERE us.user_id=$1`, usersSegmentsTable, segmentsTable)
	err := repo.db.Select(&slugs, query, userId)
	return slugs, err
}

func printSliceByComma[T uint | string](slice []T) string {
	var buf strings.Builder
	if len(slice) == 0 {
		return "''"
	}
	buf.WriteString(fmt.Sprintf("'%v'", slice[0]))
	for _, el := range slice[1:] {
		buf.WriteString(fmt.Sprintf(", '%v'", el))
	}
	return buf.String()
}

func (repo *UsersSegmentsRepository) filterSegments(slugs []string) ([]SegmentEntity, error) {
	slugsString := printSliceByComma(slugs)
	querySegmentsToAdd := fmt.Sprintf(
		`SELECT * FROM %s AS s 
            WHERE s.slug IN (%s)`, segmentsTable, slugsString)
	var segments []SegmentEntity
	err := repo.db.Select(&segments, querySegmentsToAdd)
	return segments, err
}

func (repo *UsersSegmentsRepository) insertSegmentsIntoUser(segmentsToAdd []SegmentEntity, userId uint) error {
	var usersSegmentsBuilder strings.Builder
	usersSegmentsBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_id) VALUES ",
		usersSegmentsTable))

	for i, seg := range segmentsToAdd {
		if i != 0 {
			usersSegmentsBuilder.WriteString(", ")
		}
		usersSegmentsBuilder.WriteString(
			fmt.Sprintf("('%d', '%d')", userId, seg.Id))
	}

	query := usersSegmentsBuilder.String()
	_, err := repo.db.Exec(query)
	return err
}

func (repo *UsersSegmentsRepository) deleteSegmentsFromUser(segmentsToRemove []SegmentEntity, userId uint) error {
	segmentsToRemoveIds, _ := decomposeSegments(segmentsToRemove)
	segmentsToRemoveString := printSliceByComma(segmentsToRemoveIds)
	query := fmt.Sprintf("DELETE FROM %s AS us WHERE us.segment_id IN (%s) AND us.user_id=$1", usersSegmentsTable, segmentsToRemoveString)
	_, err := repo.db.Exec(query, userId)
	return err
}

func (repo *UsersSegmentsRepository) saveInHistory(segments []SegmentEntity, userId uint, op operation) error {
	var usersSegmentsHistoryBuilder strings.Builder
	usersSegmentsHistoryBuilder.WriteString(fmt.Sprintf(
		"INSERT INTO %s (user_id, segment_slug, operation, updated_at) VALUES ",
		usersSegmentsHistoryTable))
	for i, seg := range segments {
		if i != 0 {
			usersSegmentsHistoryBuilder.WriteString(", ")
		}
		usersSegmentsHistoryBuilder.WriteString(
			fmt.Sprintf("('%d', '%s', '%s', 'now()')", userId, seg.Slug, op))
	}

	query := usersSegmentsHistoryBuilder.String()
	_, err := repo.db.Exec(query)
	return err
}
