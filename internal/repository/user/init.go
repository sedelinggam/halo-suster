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
	GetUserByNIPWithRole(ctx context.Context, nip string, role string) (*entity.User, error)
	UpdateDeletedAt(ctx context.Context, data entity.User) error
	UpdatePassword(ctx context.Context, data entity.User) error
}

func New(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
