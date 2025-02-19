package response

import "time"

type RespProvince struct {
	Id        string     `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	IsActive  bool       `json:"is_active" db:"is_active"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
