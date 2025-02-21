package repository

import (
	"context"
	"zayyid-go/domain/testimoni/model"
	"zayyid-go/infrastructure/database"
)

type TestimoniRepository interface {
	AddTestimoniRepository(ctx context.Context, request model.Testimoni) (err error)
	UpdateTestimoniRepository(ctx context.Context, request model.Testimoni) (err error)
	GetTestimoniRepository(ctx context.Context, request model.Testimoni) (response model.Testimoni, err error)
	GetListTestimoniRepository(ctx context.Context, request model.Testimoni) (response []model.Testimoni, err error)
}

type testimoniRepository struct {
	database *database.Database
}

func NewTestimoniRepository(db *database.Database) TestimoniRepository {
	return &testimoniRepository{
		database: db,
	}
}
