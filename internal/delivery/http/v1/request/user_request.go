package request

type UserRegister struct {
	NIP      int    `json:"nip" validate:"required,min=10,max=16"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type UserLogin struct {
	NIP      int    `json:"nip" validate:"required,min=10,max=16"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	RoleType string
}

type UserParam struct {
	UserID    *string
	Limit     *int
	Offset    *int
	Name      *string
	Nip       *string
	Role      *string
	CreatedAt *string
}
