package model

type TokenRes struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRes struct {
	Id             string `json:"id" db:"id"`
	UserName       string `json:"username" db:"username"`
	Name           string `json:"name" db:"name"`
	WhatsAppNumber string `json:"whatsapp_number" db:"whatsapp_number"`
	Email          string `json:"email" db:"email"`
	Role           string `json:"role" db:"role"`
	Password       string `json:"-" db:"password"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	CreatedBy      string `json:"created_by" db:"created_by"`

	// token response
	TokenData TokenRes `json:"token_data" db:"-"`
}
