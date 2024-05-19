package patientService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	patientRepository "halo-suster/internal/repository/patient"

	"github.com/jmoiron/sqlx"
)

type patientService struct {
	patientRepo patientRepository.PatientRepository
}

type PatientService interface {
	GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error)
	CreatePatient(ctx context.Context, requestData request.CreatePatient) (*response.CreatePatient, error)
}

func New(db *sqlx.DB) PatientService {
	return &patientService{
		patientRepo: patientRepository.New(db),
	}
}
