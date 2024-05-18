package medicalService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/lumen"
)

func (ms medicalService) GetMedicalRecords(ctx context.Context, req request.GetMedicalRecords) ([]*response.GetMedicalRecords, error) {
	medicalRecords, err := ms.medicalRepo.GetMedicalRecords(ctx, req)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	medicalRecordsResp := response.MapMedicalRecordsEntityToListResponse(medicalRecords)
	return medicalRecordsResp, nil
}
