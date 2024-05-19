package response

import (
	"halo-suster/internal/entity"
	"strconv"
	"time"
)

type UserAccessToken struct {
	UserID      string `json:"userId"`
	NIP         int    `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type UserNurse struct {
	UserID    string `json:"userId"`
	Nip       *int   `json:"nip,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

func MapUserEntityToResponse(e *entity.User) *UserNurse {
	nipInt, _ := strconv.Atoi(e.NIP)
	return &UserNurse{
		e.ID,
		&nipInt,
		e.Name,
		e.CreatedAt.Format(time.RFC3339),
	}
}

func MapUserListEntityToListResponse(e []*entity.User) []*UserNurse {
	var resp []*UserNurse
	for _, v := range e {
		resp = append(resp, MapUserEntityToResponse(v))
	}
	return resp
}
