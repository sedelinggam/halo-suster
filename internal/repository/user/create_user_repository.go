package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) Create(ctx context.Context, data entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, phone_number, name, password, created_at) VALUES (:id, :phone_number, :name, :password, :created_at)`, data.TableName())

	tx := sr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
