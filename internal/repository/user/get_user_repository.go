package userRepository

import (
	"context"
	"fmt"
	"halo-suster/internal/entity"
)

func (sr userRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	var (
		resp entity.User
		err  error
	)

	query := fmt.Sprintf(`SELECT * FROM %s WHERE "phone_number" = $1`, resp.TableName())

	err = sr.db.GetContext(ctx, &resp, query, phoneNumber)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
