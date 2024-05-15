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
	RegisterUserIT(ctx context.Context, requestData request.ITUserRegister) (*response.UserAccessToken, error)
	RegisterUserNurse(ctx context.Context, requestData request.NurseUserRegister) (*response.UserAccessToken, error)
	DeleteUserNurse(ctx context.Context, nip string) (*response.UserNurse, error)
	AccessUserNurse(ctx context.Context, requestData string, password string) (*response.UserNurse, error)
	UpdateUserNurse(ctx context.Context, nip int, name string, userId string) (*response.UserNurse, error)
	GetUsers(ctx context.Context, req request.UserParam) ([]*response.UserNurse, error)
}

func New(db *sqlx.DB) UserService {
	return &userService{
		userRepo: userRepository.New(db),
	}
}
