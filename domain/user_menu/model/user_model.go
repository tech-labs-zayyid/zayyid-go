package model

type User struct {
	Id               string `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	Role             string `db:"role" json:"role"`
	CreatedAt        string `db:"created_at" json:"created_at"`
	IsActive         bool   `db:"is_active" json:"is_active"`
	TempIsActive     int    `db:"temp_is_active" json:"temp_is_active"`
	CompanyPermision []byte `db:"company_permission" json:"company_permission"`
	MenuId           []int  `json:"menu_id"`

	StatusCode int    `json:"status_code"`
	Limit      int64  `json:"limit"`
	Page       int64  `json:"page"`
	Search     string `json:"search"`
}

type UserAuth struct {
	Id        int64  `db:"id" json:"id"`
	UserId    string `db:"user_id" json:"email"`
	MenuId    int    `db:"menu_id" json:"menu_id"`
	Permision []byte `db:"permission" json:"permission"`
}
