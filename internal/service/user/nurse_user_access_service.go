package userService

import (
	"context"
	"errors"
	"fmt"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/crypto/bcrypt"
	"halo-suster/package/lumen"
)

func (ss userService) AccessUserNurse(ctx context.Context, requestData int, password string) (*response.UserNurse, error) {

	//Password Hash
	hashPassword, err := bcrypt.HashPassword(password)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	userData := entity.User{
		ID:       fmt.Sprintf("%d", requestData),
		Password: hashPassword,
	}
	//Check NIP
	if validNIP := userData.CheckNIP(false); !validNIP {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("NIP not valid"))
	}

	err = ss.userRepo.UpdatePassword(ctx, userData)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	return &response.UserNurse{
		NIP: requestData,
	}, nil
}
