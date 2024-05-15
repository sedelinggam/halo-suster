package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) Create(ctx context.Context, data entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, nip, name, password, created_at, role) VALUES (:id, :nip, :name, :password, :created_at, :role)`, data.TableName())
	tx := sr.db.MustBegin()
	_, err := tx.NamedExecContext(ctx, query, data)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
