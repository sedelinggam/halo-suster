package entity

import (
	"time"
)

type Patient struct {
	ID                    string    `db:"id"`
	IdentityNumber        string    `db:"identity_number"`
	Name                  string    `db:"name"`
	PhoneNumber           string    `db:"phone_number"`
	BirthDate             string    `db:"birth_date"`
	Gender                *string   `db:"gender"`
	IdentityCardScanImage string    `db:"identity_card_scan_img"`
	CreatedAt             time.Time `db:"created_at"`
}

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
