package patientRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr patientRepository) Create(ctx context.Context, data entity.Patient) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, identity_number, phone_number, name, birth_date, gender, identity_card_scan_url, created_at) VALUES (:id, :identity_number, :phone_number, :name, :birth_date, :gender, :identity_card_scan_url, :created_at)`, data.TableName())
	tx := sr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
