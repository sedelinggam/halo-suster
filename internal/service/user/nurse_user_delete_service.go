package userService

import (
	"context"
	"fmt"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/lumen"
	"time"

	"github.com/lib/pq"
)

func (ss userService) DeleteUserNurse(ctx context.Context, requestData int) (*response.UserNurse, error) {

	err := ss.userRepo.UpdateDeletedAt(ctx, entity.User{
		ID: fmt.Sprintf("%d", requestData),
		DeletedAt: pq.NullTime{
			Time:  time.Now(),
			Valid: false,
		},
	})
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UserNurse{
		NIP: requestData,
	}, nil
}
