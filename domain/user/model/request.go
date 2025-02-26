package model

type RegisterRequest struct {
	Name string `json:"name"` // fullname 
	UserName string `json:"username"`
	WhatsappNumber string `json:"whatsapp_number"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role" validate:"required,oneof=sales agent"`
}