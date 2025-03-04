package request

type Testimoni struct {
	Id           string `json:"id" query:"id" db:"id"`
	PublicAccess string `json:"public_access" query:"public_access" db:"public_access"`
	FullName     string `json:"full_name" query:"full_name" db:"full_name"`
	Description  string `json:"description" query:"description" db:"description"`
	PhotoUrl     string `json:"photo_url" query:"photo_url" db:"photo_url"`
	IsActive     int    `json:"is_active" query:"is_active" db:"is_active"`
	CreatedAt    string `json:"created_at" query:"created_at" db:"created_at"`
	ModifiedAt   string `json:"modified_at" query:"modified_at" db:"modified_at"`
}

type TestimoniSearch struct {
	StatusCode int    `json:"status_code"`
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Search     string `json:"search"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
}
