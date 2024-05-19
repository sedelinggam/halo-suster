package medicalRepository

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/entity"

	"github.com/jmoiron/sqlx"
)

type medicalRepository struct {
	db *sqlx.DB
}

type MedicalRepository interface {
	CreateMedicalRecord(ctx context.Context, data entity.MedicalRecord) error
	GetMedicalRecords(ctx context.Context, request request.GetMedicalRecords) ([]*entity.MedicalRecords, error)
}

func New(db *sqlx.DB) MedicalRepository {
	return &medicalRepository{
		db: db,
	}
}
