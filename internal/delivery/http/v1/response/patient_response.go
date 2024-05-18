package response

type CreatePatient struct {
	ID             string `json:"id"`
	IdentityNumber string `json:"identityNumber"`
}
