package patientRepository

import (
	"context"
	"halo-suster/internal/entity"

	"github.com/jmoiron/sqlx"
)

type patientRepository struct {
	db *sqlx.DB
}

type PatientRepository interface {
	Create(ctx context.Context, data entity.Patient) error
	GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error)
}

func New(db *sqlx.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}
