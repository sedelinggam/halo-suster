package request

type CreateMedicalRecord struct {
	IdentityNumber int64  `json:"identityNumber" validate:"required"`
	Symptoms       string `json:"symptoms" validate:"required,gte=1,lte=2000"`
	Medications    string `json:"medications" validate:"required,gte=1,lte=2000"`
}
