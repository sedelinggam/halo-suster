package request

type CreatePatient struct {
	// should be 16 characters
	IdentityNumber int64 `json:"identityNumber" validate:"required"`

	/**
	- not null,
	- should be string,
	- starts with `+62`,
	- minLength 10,
	- maxLength 15
	*/
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=15"`

	// not null, should be string, minLength 3, maxLength 30
	Name string `json:"name" validate:"required,min=3,max=30"`

	// not null, should be string with ISO 8601 format
	BirthDate string `json:"birthDate" validate:"required"`

	// not null, should be enum of 'male'|'female'
	Gender string `json:"gender" validate:"required,oneof=male female"`

	// not null, should be an image url
	IdentityCardScanImg string `json:"identityCardScanImg" validate:"required,url"`
}

type PatientParam struct {
	IdentityNumber *string
	Limit          int
	Offset         int
	Name           *string
	PhoneNumber    *string
	CreatedAt      *string
}
