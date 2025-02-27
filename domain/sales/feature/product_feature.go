package feature

import (
	"context"
	"net/http"
	"zayyid-go/domain/sales/model/request"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedError "zayyid-go/domain/shared/helper/error"
)

func (f SalesFeature) AddProductSales(ctx context.Context, param request.AddProductReq) (err error) {
	tx := f.repo.OpenTransaction(ctx)

	defer func() {
		if err != nil {
			errRollback := f.repo.RollbackTransaction(tx)
			if errRollback != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errRollback)
			}
		} else {
			errCommit := f.repo.CommitTransaction(tx)
			if errCommit != nil {
				err = sharedError.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), errCommit)
			}
		}
	}()

	err = sharedHelper.Validate(param)
	if err != nil {
		return
	}

	return
}
