package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/master/model/response"
	sharedModel "zayyid-go/domain/shared/model"
	"zayyid-go/infrastructure/database"
)

type MasterRepository interface {
	GetMasterProvince(ctx context.Context, filter sharedModel.QueryRequest) (resp []response.RespProvince, err error)
	CountMasterProvince(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error)
	GetMasterCity(ctx context.Context, filter sharedModel.QueryRequest) (resp []response.RespCity, err error)
	CountMasterCity(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error)

	//transaction schema DB
	OpenTransaction(ctx context.Context) (tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx) (rollBack error)
	CommitTransaction(tx *sql.Tx) (commit error)
}

type masterRepository struct {
	database *database.Database
}

func NewMasterRepository(db *database.Database) MasterRepository {
	return &masterRepository{
		database: db,
	}
}
