package repository

import (
	"context"
	"database/sql"
	"zayyid-go/domain/user/model"
	"zayyid-go/infrastructure/database"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	database *database.Database
}

func NewUserRepository(db *database.Database) UserRepository {
	return UserRepository{
		database: db,
	}
}

type IUserRepository interface {
	RegisterRepository(ctx context.Context, payload model.RegisterRequest, userId string) (err error)
	GetUserById(ctx context.Context, userId string) (resp model.UserRes, err error)
	GetByQueryRepository(ctx context.Context, q model.QueryUser) (user model.UserRes, err error)
	GetDataUserByUserId(ctx context.Context, userId string) (resp model.UserDataResp, err error)
	UpdateRepository(ctx context.Context, payload model.UpdateUser, userId string) (err error)
	CheckExistsUserId(ctx context.Context, userId string) (exists bool, err error)
	CheckExistsSubdomain(ctx context.Context, subdomain string) (exists bool, err error)
	CheckExistsCodeReferal(ctx context.Context, referal string) (exists bool, err error)
	GetDataAgentByReferralCode(ctx context.Context, referralCode string) (resp model.UserDataResp, err error)
	GetAgentRepository(ctx context.Context, q model.QueryAgentList, userId string) (resp []model.UserRes, err error)

	RegisterRepositoryTransaction(ctx context.Context, payload model.RegisterRequest, userId string, tx *sqlx.Tx) (err error)
	MappingSalesAgent(ctx context.Context, salesId, agentId, createdBy string, trx *sqlx.Tx) (err error)

	// Transaction repo
	OpenTransaction(ctx context.Context) (tx *sql.Tx)
	RollbackTransaction(tx *sql.Tx) (rollBack error)
	CommitTransaction(tx *sql.Tx) (commit error)
}
