package userService

import (
	"context"
	"errors"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/package/crypto/bcrypt"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"strconv"
)

func (ss userService) Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error) {
	//Password Hash
	var (
		err error
	)

	// Find the user by credentials
	user, err := ss.userRepo.GetUserByNIPWithRole(ctx, strconv.Itoa(requestData.NIP), requestData.RoleType)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLNotFound(err) {
			return nil, lumen.NewError(lumen.ErrNotFound, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Compare password hash
	if !bcrypt.CheckPasswordHash(requestData.Password, user.Password) {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("password doesn't match"))
	}
	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(user.ID, user.NIP, user.UserRole)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	respAccessToken := &response.UserAccessToken{
		NIP:         requestData.NIP,
		Name:        user.Name,
		AccessToken: *accessToken,
	}

	return respAccessToken, nil
}
