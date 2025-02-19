package response

import (
	"time"
)

type RespCity struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	ProvinceId   string     `json:"province_id"`
	ProvinceName string     `json:"province_name"`
	IsActive     bool       `json:"is_active"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
