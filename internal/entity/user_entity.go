package entity

import (
	"errors"
	"halo-suster/package/lumen"
	"strings"
	"time"
)

type User struct {
	ID          string    `db:"id"`
	PhoneNumber string    `db:"phone_number"`
	Name        string    `db:"name"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
}

func (s User) TableName() string {
	return `users`
}

func (s User) CheckPhoneNumber() error {
	if len(s.PhoneNumber) == 0 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	if !strings.HasPrefix(s.PhoneNumber, "+") {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	return nil
}

func (s *User) NewPassword(password string) error {
	return nil
}
