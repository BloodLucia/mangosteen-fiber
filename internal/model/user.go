package model

type User struct {
	ID    uint64
	Email string
}

type UserInReq struct {
	VerificationCode string `json:"verificationCode"`
	Email            string `json:"email"`
}

func (u User) TableName() string {
	return "t_users"
}

func (r *UserInReq) ToModel() *User {
	return &User{
		Email: r.Email,
	}
}
