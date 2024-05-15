package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) GetUserByNIPWithRole(ctx context.Context, nip string, role string) (*entity.User, error) {
	var (
		resp entity.User
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "nip" = $1 and "role" = $2`, resp.TableName())

	err = sr.db.GetContext(ctx, &resp, query, nip, role)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
