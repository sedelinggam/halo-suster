package medicalService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	medicalRepository "halo-suster/internal/repository/medical"
	patientRepository "halo-suster/internal/repository/patient"

	"github.com/jmoiron/sqlx"
)

type medicalService struct {
	medicalRepo medicalRepository.MedicalRepository
	patientRepo patientRepository.PatientRepository
}

type MedicalService interface {
	CreateMedicalRecord(ctx context.Context, requestData request.CreateMedicalRecord) (*response.CreateMedicalRecord, error)
}

func New(db *sqlx.DB) MedicalService {
	return &medicalService{
		patientRepo: patientRepository.New(db),
		medicalRepo: medicalRepository.New(db),
	}
}
