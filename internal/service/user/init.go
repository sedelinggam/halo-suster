package userService

import (
	"context"
	"halo-suster/internal/delivery/http/v1/request"
	"halo-suster/internal/delivery/http/v1/response"
	userRepository "halo-suster/internal/repository/user"

	"github.com/jmoiron/sqlx"
)

type userService struct {
	userRepo userRepository.UserRepository
}

type UserService interface {
	Login(ctx context.Context, requestData request.UserLogin) (*response.UserAccessToken, error)
	Register(ctx context.Context, requestData request.UserRegister) (*response.UserAccessToken, error)
}

func New(db *sqlx.DB) UserService {
	return &userService{
		userRepo: userRepository.New(db),
	}
}
