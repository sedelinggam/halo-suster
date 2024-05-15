package response

type UserAccessToken struct {
	UserID      string `json:"userId"`
	NIP         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type UserNurse struct {
	NIP  int    `json:"nip"`
	Name string `json:"name,omitempty"`
}
