package response

import (
	"halo-suster/internal/entity"
	"strconv"
	"time"
)

type CreateMedicalRecord struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type PatientDetail struct {
	IdentityNumber      int       `json:"identityNumber"`
	PhoneNumber         string    `json:"phoneNumber"`
	Name                string    `json:"name"`
	BirthDate           time.Time `json:"birthDate"`
	Gender              string    `json:"gender"`
	IdentityCardScanImg string    `json:"identityCardScanImg"`
}

type CreatedBy struct {
	NIP    int    `json:"nip"`
	Name   string `json:"name"`
	UserID string `json:"userId"`
}

type GetMedicalRecords struct {
	IdentityDetail PatientDetail `json:"identityDetail"`
	Symptoms       string        `json:"symptoms"`
	Medications    string        `json:"medications"`
	CreatedAt      string        `json:"createdAt"`
	CreatedBy      CreatedBy     `json:"createdBy"`
}

func MapMedicalRecordsEntityToResponse(e *entity.MedicalRecords) *GetMedicalRecords {
	idNumber, _ := strconv.Atoi(e.IdentityNumber)
	userNIP, _ := strconv.Atoi(e.UserNIP)
	return &GetMedicalRecords{
		IdentityDetail: PatientDetail{
			IdentityNumber:      idNumber,
			PhoneNumber:         e.PhoneNumber,
			Name:                e.PatientName,
			BirthDate:           e.BirthDate,
			Gender:              e.Gender,
			IdentityCardScanImg: e.IdentityCardScanImg,
		},
		Symptoms:    e.Symptoms,
		Medications: e.Medications,
		CreatedAt:   e.CreatedAt.Format("2006-01-02T15:04:05.999999Z07:00"),
		CreatedBy: CreatedBy{
			NIP:    userNIP,
			Name:   e.UserName,
			UserID: e.UserID,
		},
	}
}

func MapMedicalRecordsEntityToListResponse(e []*entity.MedicalRecords) []*GetMedicalRecords {
	var resp []*GetMedicalRecords
	for _, v := range e {
		resp = append(resp, MapMedicalRecordsEntityToResponse(v))
	}
	return resp
}
