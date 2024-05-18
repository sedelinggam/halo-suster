package medicalService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

func (ms medicalService) CreateMedicalRecord(ctx context.Context, requestData request.CreateMedicalRecord) (*response.CreateMedicalRecord, error) {
	var (
		err error
	)

	idNumber := strconv.FormatInt(requestData.IdentityNumber, 10)
	_, err = ms.patientRepo.GetPatient(ctx, idNumber)

	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	currentUser := ctx.Value("currentUser").(*cryptoJWT.JWTClaims)

	newMedicalRecord := entity.MedicalRecord{
		ID:             ulid.Make().String(),
		CreatedAt:      time.Now(),
		Symptoms:       requestData.Symptoms,
		Medications:    requestData.Medications,
		IdentityNumber: idNumber,
		UserID:         currentUser.Id,
	}

	err = ms.medicalRepo.CreateMedicalRecord(ctx, newMedicalRecord)

	if err != nil {
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	resp := &response.CreateMedicalRecord{
		ID:        newMedicalRecord.ID,
		CreatedAt: newMedicalRecord.CreatedAt.Format(time.RFC3339),
	}
	return resp, nil
}
