package request

type CreateMedicalRecord struct {
	IdentityNumber string `json:"identityNumber" validate:"required,len=16"`
	Symptoms       string `json:"symptoms" validate:"required,gte=1,lte=2000"`
	Medications    string `json:"medications" validate:"required,gte=1,lte=2000"`
}
