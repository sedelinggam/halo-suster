package userRepository

import (
	"context"
	"errors"
	"fmt"
	"halo-suster/internal/entity"
)

func (cr userRepository) UpdatePassword(ctx context.Context, data entity.User) error {
	query := fmt.Sprintf(`UPDATE %s SET password = $1 WHERE id = $2 AND role = 'nurse'`, data.TableName())

	tx := cr.db.MustBegin()
	res, err := tx.ExecContext(ctx, query, data.Password, data.ID)
	tx.Commit()
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return errors.New("no rows in result set")
	}

	return nil
}
