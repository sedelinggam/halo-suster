package patientRepository

import (
	"context"
	"halo-suster/internal/entity"
)

func (pr patientRepository) GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error) {
	var (
		data entity.Patient
	)

	err := pr.db.Get(&data, `SELECT * FROM patients WHERE identity_number=$1`, identityNumber)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
