package patientService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
)

func (cs patientService) GetPatients(ctx context.Context, requestData request.PatientParam) ([]*response.Patient, error) {
	customer, err := cs.patientRepo.GetPatients(ctx, requestData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return response.MapPatientListEntityToListResponse(customer), nil
}
