package model

import (
	"time"
	sharedModel "zayyid-go/domain/shared/model"
)

type UserMenuResponse struct {
	Body       []byte
	StatusCode int // Original status from target server
	Status     string
	Duration   time.Duration
}

type UserResponse struct {
	Id               string   `db:"id" json:"id"`
	Name             string   `db:"name" json:"name"`
	Role             string   `db:"role" json:"role"`
	CreatedAt        string   `db:"created_at" json:"created_at"`
	IsActive         bool     `db:"is_active" json:"is_active"`
	TempIsActive     int      `db:"temp_is_active" json:"temp_is_active"`
	CompanyPermision []string `db:"company_permission" json:"company_permission"`
	MenuId           []int    `json:"menu_id"`

	StatusCode int `json:"code"`
}

type Menu struct {
	Id          int64  `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Permision   []byte `db:"permission" json:"permission"`
}

type GetListMenu struct {
	Data       []Menu `json:"data"`
	StatusCode int    `json:"code"`
}

type GetList struct {
	Data       []User                 `json:"data"`
	Pagination sharedModel.Pagination `json:"pagination"`
	StatusCode int                    `json:"code"`
}

type ParamUserResponse struct {
	Id         string `json:"id" db:"id"`
	StatusCode int    `json:"status_code"`
}

type AppType struct {
	Data       []string `json:"data"`
	StatusCode int      `json:"status_code"`
}
