package entity

import (
	"errors"
	"halo-suster/package/lumen"
	"time"
)

type MedicalRecord struct {
	ID             string    `db:"id"`
	CreatedAt      time.Time `db:"created_at"`
	Symptoms       string    `db:"symptoms"`
	Medications    string    `db:"medications"`
	IdentityNumber string    `db:"identity_number"`
	UserID         string    `db:"user_id"`
}

func (m MedicalRecord) TableName() string {
	return `medical_records`
}

func (m MedicalRecord) CheckIdentityNumber() error {
	if len(m.IdentityNumber) != 16 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("identity number not valid"))
	}
	return nil
}
