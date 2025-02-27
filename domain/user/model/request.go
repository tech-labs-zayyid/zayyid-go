package model

type RegisterRequest struct {
	Name           string `json:"name"` // fullname
	UserName       string `json:"username"`
	WhatsappNumber string `json:"whatsapp_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role" validate:"required,oneof=sales agent"`
}

type AuthUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type QueryUser struct {
	Id    string
	Email string
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
