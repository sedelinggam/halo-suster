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

type MedicalRecords struct {
	MedicalRecord
	PhoneNumber         string    `db:"phone_number"`
	PatientName         string    `db:"patient_name"`
	BirthDate           time.Time `db:"birth_date"`
	Gender              string    `db:"gender"`
	IdentityCardScanImg string    `db:"identity_card_scan_url"`
	UserNIP             string    `db:"nip"`
	UserName            string    `db:"user_name"`
}
