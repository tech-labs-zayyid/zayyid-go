package model

type GetListRequest struct {
	CompanyId string `query:"company_id"`

	// Pagination
	Page   int64 `query:"page"`
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

type ApplicationAPIConfigRow struct {
	BasePath string `db:"base_path"`
}

type ApplicationSAPConfigRow struct {
	SAPClient     int    `db:"sap_client"`
	SAPProgram    string `db:"sap_program"`
	TargetBaseUrl string `db:"target_base_url"`
}
