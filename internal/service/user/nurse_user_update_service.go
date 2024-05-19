package userService

import (
	"context"
	"errors"
	"fmt"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/lumen"
)

func (ss userService) UpdateUserNurse(ctx context.Context, nip int, name string, userId string) (*response.UserNurse, error) {

	userData := entity.User{
		ID:   userId,
		Name: name,
		NIP:  fmt.Sprintf("%d", nip),
	}
	//Check NIP
	if validNIP := userData.CheckNIP(false); !validNIP {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("NIP not valid"))
	}

	err := ss.userRepo.UpdateUser(ctx, userData)
	if err != nil {
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		} else if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UserNurse{
		UserID: userId,
	}, nil
}
