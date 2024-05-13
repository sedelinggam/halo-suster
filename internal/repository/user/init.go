package userRepository

import (
	"context"
	"halo-suster/internal/entity"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, data entity.User) error
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
