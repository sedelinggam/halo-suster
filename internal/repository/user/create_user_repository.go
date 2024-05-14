package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) Create(ctx context.Context, data entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s(id, nip, name, password, created_at, role) VALUES (?, ?, ?, ?, ?, ?)`, data.TableName())

	tx := sr.db.MustBegin()
	_, err := tx.ExecContext(ctx, query, data.ID, data.NIP, data.Name, data.Password, data.CreatedAt, data.UserRole)
	tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
