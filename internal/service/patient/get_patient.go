package patientService

import (
	"context"
	"halo-suster/internal/entity"
	"halo-suster/package/lumen"
)

func (ps patientService) GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error) {
	var (
		err error
	)

	data, err := ps.patientRepo.GetPatient(ctx, identityNumber)

	if err != nil {
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return data, nil
}
