package feature

import (
	"context"
	"errors"
	"github.com/gosimple/slug"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"zayyid-go/domain/sales/helper"
	"zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	sharedContext "zayyid-go/domain/shared/context"
	sharedHelper "zayyid-go/domain/shared/helper"
	sharedConstant "zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	paginate "zayyid-go/domain/shared/helper/pagination"
	sharedModel "zayyid-go/domain/shared/model"
)

func (f salesFeature) AddProductSales(ctx context.Context, param request.AddProductReq) (err error) {
	valueCtx := sharedContext.GetValueContext(ctx)
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

	if len(param.Images) == 0 || len(param.Images) > 5 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestProductImages, err)
		return
	}

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	re := regexp.MustCompile(`\W`)
	cleaned := re.ReplaceAllString(param.ProductName, "")
	existsProductName, err := f.repo.CheckExistsProductName(ctx, cleaned, valueCtx.UserId)
	if err != nil {
		return
	}

	if existsProductName {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDuplicateProductName, errors.New(sharedConstant.ErrDuplicateProductName))
		return
	}

	dataUser, err := f.userRepo.GetDataUserByUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	param.ProductCategoryName = helper.CARS_SALES_PRODUCT_CATEGORY_PAGE.PageCategory()
	param.Status = helper.PRODUCT_LISTED.StatusProduct()
	param.SalesId = valueCtx.UserId
	param.PublicAccess = dataUser.Username
	param.Slug = slug.Make(param.ProductName)
	err = f.repo.AddProductSales(ctx, tx, param)
	return
}

func (f salesFeature) ListProductSales(ctx context.Context, paramFilter sharedModel.QueryRequest) (resp []response.ProductListBySales, pagination *sharedModel.Pagination, err error) {
	valueCtx := sharedContext.GetValueContext(ctx)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	paramFilter.SalesId = valueCtx.UserId
	respData, err := f.repo.GetListProduct(ctx, paramFilter)
	if err != nil {
		return
	}

	for _, v := range respData {
		resp = append(resp, *v)
	}

	count, err := f.repo.CountListProduct(ctx, paramFilter)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, paramFilter.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = paramFilter.Page
	return
}

func (f salesFeature) GetDetailSalesProduct(ctx context.Context, id string) (resp response.ProductDetailResp, err error) {
	valueCtx := sharedContext.GetValueContext(ctx)

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existsProduct, err := f.repo.CheckExistsProductId(ctx, id)
	if err != nil {
		return
	}

	if !existsProduct {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdProductNotFound, errors.New(sharedConstant.ErrIdProductNotFound))
		return
	}

	resp, err = f.repo.DetailSalesProduct(ctx, id)
	return
}

func (f salesFeature) UpdateProductSales(ctx context.Context, param request.UpdateProductSales) (err error) {
	valueCtx := sharedContext.GetValueContext(ctx)
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

	validateStatus := helper.StatusProductStr(strings.ToLower(param.Status)).IsValid()
	if !validateStatus {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrStatusInvalid, errors.New(sharedConstant.ErrStatusInvalid))
		return
	}

	exists, err := f.userRepo.CheckExistsUserId(ctx, valueCtx.UserId)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	existsProduct, err := f.repo.CheckExistsProductId(ctx, param.ProductId)
	if err != nil {
		return
	}

	if !existsProduct {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrIdProductNotFound, errors.New(sharedConstant.ErrIdProductNotFound))
		return
	}

	count, err := f.repo.GetCountDataImageByProductId(ctx, param.ProductId)
	if err != nil {
		return
	}

	if len(param.Images) == 0 || len(param.Images) > 5 || count+len(param.Images) > 5 {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrRequestProductImages, err)
		return
	}

	dataProduct, err := f.GetDetailSalesProduct(ctx, param.ProductId)
	if err != nil {
		return
	}

	if dataProduct.ProductName != param.ProductName {
		re := regexp.MustCompile(`\W`)
		cleaned := re.ReplaceAllString(param.ProductName, "")
		existsProductName, errExistProductName := f.repo.CheckExistsProductName(ctx, cleaned, valueCtx.UserId)
		if errExistProductName != nil {
			err = errExistProductName
			return
		}

		if existsProductName {
			err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDuplicateProductName, errors.New(sharedConstant.ErrDuplicateProductName))
			return
		}
	}

	if dataProduct.Status != param.Status {
		if err = f.repo.ChangeStatusProductSales(ctx, tx, param); err != nil {
			return
		}
	}

	param.Slug = slug.Make(param.ProductName)
	err = f.repo.UpdateProductSales(ctx, tx, param)
	return
}

func (f salesFeature) GetListProductSalesPublic(ctx context.Context, filter request.ProductListPublic, subdomain string) (resp []response.ProductListSalesPublic, pagination *sharedModel.Pagination, err error) {
	var (
		minimumPrice float64
		maximumPrice float64
	)

	exists, err := f.userRepo.CheckExistsSubdomain(ctx, subdomain)
	if err != nil {
		return
	}

	if !exists {
		err = sharedError.New(http.StatusBadRequest, sharedConstant.ErrDataUserIdNotFound, errors.New(sharedConstant.ErrDataUserIdNotFound))
		return
	}

	sharedHelper.SetDefaults(&filter)
	if filter.MinimumPrice != "" {
		minimumPrice, err = strconv.ParseFloat(filter.MinimumPrice, 32)
		if err != nil {
			return
		}
	}

	if filter.MaximumPrice != "" {
		maximumPrice, err = strconv.ParseFloat(filter.MaximumPrice, 32)
		if err != nil {
			return
		}
	}

	queryRequest := sharedModel.QueryRequest{
		Search:             filter.Search,
		SubCategoryProduct: filter.SubCategoryProduct,
		BestProduct:        filter.BestProduct,
		StatusProduct:      filter.StatusProduct,
		MinimumPrice:       minimumPrice,
		MaximumPrice:       maximumPrice,
		Page:               filter.Page,
		Limit:              filter.Limit,
		SortBy:             filter.SortBy,
		SortOrder:          filter.SortOrder,
		PublicAccess:       subdomain,
	}

	respList, err := f.repo.GetListProductSalesPublic(ctx, queryRequest)
	if err != nil {
		return
	}

	for _, v := range respList {
		resp = append(resp, *v)
	}

	count, err := f.repo.CountListProductSalesPublic(ctx, queryRequest)
	if err != nil {
		return
	}

	pagination, err = paginate.CalculatePagination(ctx, filter.Limit, count)
	if err != nil {
		return
	}

	pagination.Page = filter.Page
	return
}
