package patientService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/lumen"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

func (ps patientService) CreatePatient(ctx context.Context, requestData request.CreatePatient) (*response.CreatePatient, error) {
	var (
		err error
	)

	//Create User
	patient := entity.Patient{
		ID:                  ulid.Make().String(),
		IdentityNumber:      strconv.FormatInt(requestData.IdentityNumber, 10),
		PhoneNumber:         requestData.PhoneNumber,
		Name:                requestData.Name,
		BirthDate:           requestData.BirthDate,
		Gender:              requestData.Gender,
		IdentityCardScanUrl: requestData.IdentityCardScanImg,
		CreatedAt:           time.Now(),
	}

	err = patient.CheckIdentityNumber()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = patient.CheckPhoneNumber()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = patient.CheckBirthDate()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = ps.patientRepo.Create(ctx, patient)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Cache The token
	resp := &response.CreatePatient{
		ID:             patient.ID,
		IdentityNumber: patient.IdentityNumber,
	}
	return resp, nil
}
