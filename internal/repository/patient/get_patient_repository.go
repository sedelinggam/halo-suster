package patientRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (pr patientRepository) GetPatient(ctx context.Context, identityNumber string) (*entity.Patient, error) {
	var (
		data entity.Patient
	)

	err := pr.db.GetContext(ctx, &data, `SELECT * FROM patients WHERE identity_number=$1`, identityNumber)

	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return &data, nil
}
