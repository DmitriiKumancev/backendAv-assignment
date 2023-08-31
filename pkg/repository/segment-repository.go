package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SegmentRepository struct {
	db *sqlx.DB
}

type SegmentEntity struct {
	Id   uint
	Slug string
}

func NewSegmentRepository(db *sqlx.DB) *SegmentRepository {
	return &SegmentRepository{db: db}
}

func (repo *SegmentRepository) Create(slug string) (uint, error) {
	query := fmt.Sprintf("INSERT INTO %s (slug) VALUES ($1) RETURNING id", segmentsTable)
	var id uint
	row := repo.db.QueryRow(query, slug)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *SegmentRepository) Delete(slug string) error {
	query := fmt.Sprintf("DELETE FROM %s AS s WHERE s.slug=$1", segmentsTable)
	_, err := repo.db.Exec(query, slug)
	return err
}

func decomposeSegments(segments []SegmentEntity) ([]uint, []string) {
	var ids []uint
	var slugs []string
	for _, seg := range segments {
		ids = append(ids, seg.Id)
		slugs = append(slugs, seg.Slug)
	}
	return ids, slugs
}
