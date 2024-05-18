package medicalRepository

import (
	"context"
	"halo-suster/internal/entity"

	"github.com/jmoiron/sqlx"
)

type medicalRepository struct {
	db *sqlx.DB
}

type MedicalRepository interface {
	CreateMedicalRecord(ctx context.Context, data entity.MedicalRecord) error
}

func New(db *sqlx.DB) MedicalRepository {
	return &medicalRepository{
		db: db,
	}
}
