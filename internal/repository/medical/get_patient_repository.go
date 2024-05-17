package medicalRepository

import (
	"context"
	"halo-suster/internal/entity"
)

func (mr medicalRepository) GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error) {
	var (
		data *entity.Patient
	)

	err := mr.db.Select(&data, `SELECT * FROM patients WHERE identity_number=$1`, identityNumber)

	if err != nil {
		return nil, err
	}
	return data, nil
}
