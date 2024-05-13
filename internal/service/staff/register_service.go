package staffService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	"halo-suster/internal/entity"
	"halo-suster/package/crypto/bcrypt"
	cryptoJWT "halo-suster/package/crypto/jwt"
	"halo-suster/package/lumen"
	"time"

	"github.com/google/uuid"
)

func (ss staffService) Register(ctx context.Context, requestData request.StaffRegister) (*response.UserAccessToken, error) {
	var (
		err          error
		hashPassword string
	)

	//Password Hash
	hashPassword, err = bcrypt.HashPassword(requestData.Password)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}
	//Create User
	userData := entity.Staff{
		ID:          uuid.New().String(),
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
		Password:    hashPassword,
		CreatedAt:   time.Now(),
	}

	//Check Phone Number
	err = userData.CheckPhoneNumber()
	if err != nil {
		return nil, lumen.NewError(lumen.ErrBadRequest, err)
	}

	err = ss.staffRepo.Create(ctx, userData)
	if err != nil {
		//Duplicate unique key
		if lumen.CheckErrorSQLUnique(err) {
			return nil, lumen.NewError(lumen.ErrConflict, err)
		}
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	// Create the Claims
	accessToken, err := cryptoJWT.GenerateToken(userData.ID, userData.PhoneNumber)
	if err != nil {
		return nil, lumen.NewError(lumen.ErrInternalFailure, err)
	}

	//Cache The token
	respAccessToken := &response.UserAccessToken{
		UserID:      userData.ID,
		PhoneNumber: requestData.PhoneNumber,
		Name:        requestData.Name,
		AccessToken: *accessToken,
	}
	return respAccessToken, nil
}
