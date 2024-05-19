package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) GetUserByNIP(ctx context.Context, nip string) (*entity.User, error) {
	var (
		resp entity.User
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "nip" = $1`, resp.TableName())

	err = sr.db.Get(&resp, query, nip)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
