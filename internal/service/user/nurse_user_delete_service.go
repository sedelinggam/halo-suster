package userService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/lumen"
	"time"

	"github.com/lib/pq"
)

func (ss userService) DeleteUserNurse(ctx context.Context, requestData string) (*response.UserNurse, error) {

	err := ss.userRepo.UpdateDeletedAt(ctx, entity.User{
		ID: requestData,
		DeletedAt: pq.NullTime{
			Time:  time.Now(),
			Valid: false,
		},
	})
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UserNurse{
		UserID: requestData,
	}, nil
}
