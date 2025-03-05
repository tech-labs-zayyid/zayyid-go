package model

type RegisterRequest struct {
	Name           string `json:"name"` // fullname
	UserName       string `json:"username"`
	WhatsappNumber string `json:"whatsapp_number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ImageUrl       string `json:"image_url"`
	ReferalCode    string `json:"-"`
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

type UpdateUser struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	WhatsappNumber string `json:"whatsapp_number"`
	Password       string `json:"password"`
	ImageUrl       string `json:"image_url"`
}

type QueryAgentList struct {
	Search string `json:"search"`
	Limit int `json:"limit" validate:"required"`
	Page int `json:"page" validate:"required"`
	Sort string `json:"sort"` 
}