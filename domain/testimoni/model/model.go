package model

type Testimoni struct {
	Id         string `json:"id" query:"id" db:"id"`
	UserName   string `json:"user_name" query:"user_name" db:"user_name"`
	Position   string `json:"position" query:"position" db:"position"`
	Deskripsi  string `json:"deskripsi" query:"deskripsi" db:"deskripsi"`
	PhotoUrl   string `json:"photo_url" query:"photo_url" db:"photo_url"`
	IsActive   int    `json:"is_active" query:"is_active" db:"is_active"`
	CreatedAt  string `json:"created_at" query:"created_at" db:"created_at"`
	ModifiedAt string `json:"modified_at" query:"modified_at" db:"modified_at"`

	IsUpdate   int    `json:"is_update"`
	StatusCode int    `json:"status_code"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Search     string `json:"search"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
}

type UserAuth struct {
	Id        int64  `db:"id" json:"id"`
	UserId    string `db:"user_id" json:"email"`
	MenuId    int    `db:"menu_id" json:"menu_id"`
	Permision []byte `db:"permission" json:"permission"`
}
