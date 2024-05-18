package response

import (
	"halo-suster/internal/entity"
	"strconv"
)

type CreatePatient struct {
	ID             string `json:"id"`
	IdentityNumber string `json:"identityNumber"`
}

type Patient struct {
	IdentityNumber      int    `json:"identityNumber"`
	PhoneNumber         string `json:"phoneNumber"`
	Name                string `json:"name"`
	BirthDate           string `json:"birthDate"`
	Gender              string `json:"gender"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

func MapPatientEntityToResponse(e *entity.Patient) *Patient {
	identity, _ := strconv.Atoi(e.IdentityNumber)
	return &Patient{
		identity,
		e.PhoneNumber,
		e.Name,
		e.BirthDate,
		e.Gender,
		e.IdentityCardScanUrl,
	}
}

func MapPatientListEntityToListResponse(e []*entity.Patient) []*Patient {
	var resp []*Patient
	for _, v := range e {
		resp = append(resp, MapPatientEntityToResponse(v))
	}
	return resp
}
