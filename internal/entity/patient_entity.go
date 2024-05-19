package entity

import (
	"errors"
	"halo-suster/package/lumen"
	"regexp"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Patient struct {
	ID                  string      `db:"id"`
	IdentityNumber      string      `db:"identity_number"`
	PhoneNumber         string      `db:"phone_number"`
	Name                string      `db:"name"`
	BirthDate           string      `db:"birth_date"`
	Gender              string      `db:"gender"`
	IdentityCardScanUrl string      `db:"identity_card_scan_url"`
	CreatedAt           time.Time   `db:"created_at"`
	DeletedAt           pq.NullTime `db:"deleted_at"`
}

func (p Patient) TableName() string {
	return `patients`
}

func (p Patient) CheckPhoneNumber() error {
	if len(p.PhoneNumber) == 0 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	if !strings.HasPrefix(p.PhoneNumber, "+62") {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("phone number not valid"))
	}
	return nil
}

func (p Patient) CheckIdentityNumber() error {
	if len(p.IdentityNumber) != 16 {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("identity number not valid"))
	}
	return nil
}

func (p Patient) CheckBirthDate() error {
	const ISO8601DateRegexString = `^(-?(?:[1-9][0-9]*)?[0-9]{4})-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])(?:T|\\s)(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])?(Z)?$`
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)

	if !ISO8601DateRegex.MatchString(p.BirthDate) {
		return lumen.NewError(lumen.ErrBadRequest, errors.New("birth date not valid"))
	}

	return nil
}
