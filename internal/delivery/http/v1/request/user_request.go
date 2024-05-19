package request

type ITUserRegister struct {
	NIP      int    `json:"nip" validate:"required"`
	Name     string `json:"name" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type NurseUserRegister struct {
	NIP                   int    `json:"nip" validate:"required"`
	Name                  string `json:"name" validate:"required,min=5,max=50"`
	IdentityCardScanImage string `json:"identityCardScanImg" validate:"required"`
}

type NurseUserUpdate struct {
	NIP  int    `json:"nip" validate:"required"`
	Name string `json:"name" validate:"required,min=5,max=50"`
}

type NurseUserPassword struct {
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type UserLogin struct {
	NIP      int    `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required,min=5,max=15"`
	RoleType string
}

type UserParam struct {
	UserID    *string
	Limit     int
	Offset    int
	Name      *string
	Nip       *string
	Role      *string
	CreatedAt *string
}
