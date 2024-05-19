package medicalRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr medicalRepository) CreateMedicalRecord(ctx context.Context, data entity.MedicalRecord) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, created_at, symptoms, medications, identity_number, user_id) 
	VALUES (:id, :created_at, :symptoms, :medications, :identity_number, :user_id)`, data.TableName())
	_, err := sr.db.NamedExec(query, data)
	if err != nil {
		return err
	}
	return nil
}
