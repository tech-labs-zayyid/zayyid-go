package model

import (
	sharedModel "zayyid-go/domain/shared/model"
)

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
	ImageUrl       string `json:"image_url" db:"image_url"`
	ReferalCode    string `json:"referal_code" db:"referal_code"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	CreatedBy      string `json:"created_by" db:"created_by"`

	// token response
	TokenData TokenRes `json:"token_data" db:"-"`
}

type UserDataResp struct {
	UserId         string `json:"user_id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	WhatsappNumber string `json:"whatsapp_number"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	Desc           string `json:"desc"`
	ImageUrl       string `json:"image_url"`
}

type AgentListPagination struct {
	Data []UserDataResp `json:"docs"`
	Pagination sharedModel.Pagination
}