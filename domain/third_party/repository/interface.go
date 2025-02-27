package repository

import (
	"zayyid-go/infrastructure/database"
)

type ThirdPartyRepository struct {
	database *database.Database
}

func NewThirdPartyRepository(db *database.Database) ThirdPartyRepository {
	return ThirdPartyRepository{
		database: db,
	}
}

type ThirdPartyRepositoryInterface interface {
}
