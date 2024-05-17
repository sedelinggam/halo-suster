package medicalService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	medicalRepository "halo-suster/internal/repository/medical"

	"github.com/jmoiron/sqlx"
)

type medicalService struct {
	medicalRepo medicalRepository.MedicalRepository
}

type MedicalService interface {
	GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error)
	CreateMedicalRecord(ctx context.Context, requestData request.CreateMedicalRecord) (*response.CreateMedicalRecord, error)
}

func New(db *sqlx.DB) MedicalService {
	return &medicalService{
		medicalRepo: medicalRepository.New(db),
	}
}
