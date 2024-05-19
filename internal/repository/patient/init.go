package patientRepository

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/entity"

	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

type PatientRepository interface {
	Create(ctx context.Context, data entity.Patient) error
	GetPatients(ctx context.Context, req request.PatientParam) ([]*entity.Patient, error)
	GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error)
}

func New(db *sqlx.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}
