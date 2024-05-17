package medicalRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr medicalRepository) CreateMedicalRecord(ctx context.Context, data entity.MedicalRecord) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, created_at, symptoms, medications, identity_number, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6)`, data.TableName())
	tx := sr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
