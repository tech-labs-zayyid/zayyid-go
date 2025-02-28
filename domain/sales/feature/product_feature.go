package feature

import (
	"context"
	"errors"
	"net/http"
	"zayyid-go/domain/sales/helper"
	"zayyid-go/domain/sales/model/request"
	sharedContext "zayyid-go/domain/shared/context"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
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

	valueCtx := sharedContext.GetValueContext(ctx)

	exists, err := f.repo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	param.ProductCategoryId = helper.CarsSalesProductCategoryPage
	param.ProductCategoryName = helper.CarsSalesProductCategoryPage.PageCategory()
	param.StatusId = helper.ProductListed
	param.StatusName = helper.ProductListed.StatusProduct()
	param.SalesId = "01951f6b-db3f-7d07-8b2c-80d2e2d1be30"
	err = f.repo.AddProductSales(ctx, tx, param)
	return
}
