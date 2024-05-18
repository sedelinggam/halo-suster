package patientService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	patientRepository "halo-suster/internal/repository/patient"

	"github.com/jmoiron/sqlx"
)

type patientService struct {
	patientRepo patientRepository.PatientRepository
}

type PatientService interface {
	CreatePatient(ctx context.Context, requestData request.CreatePatient) (*response.CreatePatient, error)
	GetPatients(ctx context.Context, requestData request.PatientParam) ([]*response.Patient, error)
}

func New(db *sqlx.DB) PatientService {
	return &patientService{
		patientRepo: patientRepository.New(db),
	}
}
