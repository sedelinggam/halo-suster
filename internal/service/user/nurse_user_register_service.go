package userService

import (
	"context"
	"errors"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	valueobject "halo-suster/internal/value_object"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

func (ss userService) RegisterUserNurse(ctx context.Context, requestData request.NurseUserRegister) (*response.UserAccessToken, error) {
	var (
		err error
	)

	//Create User
	userData := entity.User{
		ID:                    ulid.Make().String(),
		NIP:                   strconv.Itoa(requestData.NIP),
		Name:                  requestData.Name,
		IdentityCardScanImage: &requestData.IdentityCardScanImage,
		CreatedAt:             time.Now(),
		UserRole:              valueobject.USER_ROLE_NURSE,
	}

	//Check NIP
	if validNIP := userData.CheckNIP(false); !validNIP {
		return nil, lumen.NewError(lumen.ErrBadRequest, errors.New("NIP not valid"))
	}

	err = ss.userRepo.Create(ctx, userData)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(userData.ID, userData.NIP, userData.UserRole)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Cache The token
	respAccessToken := &response.UserAccessToken{
		UserID:      userData.ID,
		NIP:         requestData.NIP,
		Name:        requestData.Name,
		AccessToken: *accessToken,
	}
	return respAccessToken, nil
}
