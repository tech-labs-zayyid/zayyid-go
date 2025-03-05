package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	modelRequest "zayyid-go/domain/sales/model/request"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	sharedModel "zayyid-go/domain/shared/model"
	sharedRepo "zayyid-go/domain/shared/repository"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) CheckExistsProductName(ctx context.Context, productName, salesId string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.sales_product 
        	WHERE regexp_replace(LOWER(product_name), '[^a-zA-Z0-9]', '', 'g') = $1 
        	AND sales_id = $2 AND is_active = TRUE)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, strings.ToLower(productName), salesId).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) AddProductSales(ctx context.Context, tx *sql.Tx, param modelRequest.AddProductReq) (err error) {
	var (
		id = sharedRepo.GenerateUuidAsIdTable().String()
	)

	param.ProductId = sharedRepo.GenerateUuidAsIdTable().String()
	stmtProduct, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product (id, page_category_name, sub_category_product, 
            product_name, price, tdp, installment, best_product, city_id, sales_id, public_access, slug) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductDesc, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_description (id, product_id, description) VALUES ($1, $2, $3)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductStatus, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_status (id, product_id, status) VALUES ($1, $2, $3)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductImg, err := tx.PrepareContext(ctx, `INSERT INTO product_marketing.sales_product_image (id, product_id, image_url) VALUES ($1, $2, $3)`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProduct.ExecContext(ctx, param.ProductId, param.ProductCategoryName, param.ProductSubCategory,
		param.ProductName, param.Price, param.TDP, param.Installment, param.BestProduct, param.CityId, param.SalesId, param.PublicAccess, param.Slug); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProductDesc.ExecContext(ctx, id, param.ProductId, param.Description); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProductStatus.ExecContext(ctx, id, param.ProductId, param.Status); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	for _, v := range param.Images {
		id = sharedRepo.GenerateUuidAsIdTable().String()
		if _, err = stmtProductImg.ExecContext(ctx, id, param.ProductId, v.ImageUrl); err != nil {
			err = sharedError.HandleError(err)
			return
		}
	}

	return
}

func (r salesRepository) GetProductTier(ctx context.Context) (resp response.TierResp, err error) {
	query := `SELECT pt.id, pt.tier_name, ptd.feature, ptd.limitation, ptd.length_limitation
			FROM product_marketing.product_tier pt 
			LEFT JOIN product_marketing.product_tier_detail ptd ON pt.id = ptd.tier_id
			WHERE pt.is_active = TRUE AND ptd.is_active = TRUE`

	err = r.database.QueryRowContext(ctx, query).Scan(&resp.Id, &resp.TierName, &resp.Feature, &resp.Limitation, &resp.LengthLimitation)
	if err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) GetListProduct(ctx context.Context, filter sharedModel.QueryRequest) (resp map[string]*response.ProductListBySales, err error) {
	var (
		data      response.ProductListBySales
		dataImage response.ProductImage
		dataMap   = make(map[string]*response.ProductListBySales)
		argIndex  = 1
		offset    = (filter.Page - 1) * filter.Limit
		args      = []interface{}{}
	)

	// Construct the SQL query
	query := `SELECT sp.id, COALESCE(sp.sub_category_product, ''), sp.product_name, sp.price, sp.tdp, sp.installment,
			sp.best_product, sp.is_active, spd.id, spd.description, statusProduct.status, spi.id, 
			spi.image_url, COALESCE(mp.id, '') AS province_id, COALESCE(mp.name, '') AS province_name, 
			COALESCE(mc.id, '') AS city_id, COALESCE(mc.name, '') AS city_name, sp.created_at, sp.updated_at
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE 1 = 1`

	if filter.Search != "" {
		query += fmt.Sprintf(" AND sp.product_name ILIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.SalesId != "" {
		query += fmt.Sprintf(" AND sp.sales_id = $%d", argIndex)
		args = append(args, filter.SalesId)
		argIndex++
	}

	if filter.SubCategoryProduct != "" {
		query += fmt.Sprintf(" AND sp.sub_category_product = $%d", argIndex)
		args = append(args, filter.SubCategoryProduct)
		argIndex++
	}

	if filter.StatusProduct != "" {
		query += fmt.Sprintf(" AND statusProduct.status = $%d", argIndex)
		args = append(args, filter.StatusProduct)
		argIndex++
	}

	if filter.BestProduct != "" {
		bestProduct, errParse := strconv.ParseBool(filter.BestProduct)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.best_product = $%d", argIndex)
		args = append(args, bestProduct)
		argIndex++
	}

	if filter.IsActive != "" {
		isActive, errParse := strconv.ParseBool(filter.IsActive)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.is_active = $%d", argIndex)
		args = append(args, isActive)
		argIndex++
	}

	if filter.SortBy != "" {
		query += " ORDER BY " + filter.SortBy

		if filter.SortOrder != "" {
			query += " " + filter.SortOrder
		}
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.Limit, offset)

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.IdProduct, &data.ProductSubCategory, &data.ProductName, &data.Price,
			&data.TDP, &data.Installment, &data.BestProduct, &data.IsActive, &data.IdDescription,
			&data.Description, &data.Status, &dataImage.ProductImageId, &dataImage.ImageUrl,
			&data.ProvinceId, &data.ProvinceName, &data.CityId, &data.CityName, &data.CreatedAt,
			&data.UpdatedAt); err != nil {
			err = sharedError.HandleError(err)
			return
		}

		if _, ok := dataMap[data.IdProduct]; !ok {
			dataMap[data.IdProduct] = &response.ProductListBySales{
				IdProduct:          data.IdProduct,
				ProductName:        data.ProductName,
				Price:              data.Price,
				ProductSubCategory: data.ProductSubCategory,
				TDP:                data.TDP,
				Installment:        data.Installment,
				ProvinceId:         data.ProvinceId,
				ProvinceName:       data.ProvinceName,
				CityId:             data.CityId,
				CityName:           data.CityName,
				IdDescription:      data.IdDescription,
				Description:        data.Description,
				Status:             data.Status,
				IsActive:           data.IsActive,
				ProductImages:      []response.ProductImage{},
			}
		}

		if dataImage.ProductImageId != "" {
			dataMap[data.IdProduct].ProductImages = append(dataMap[data.IdProduct].ProductImages, response.ProductImage{
				ProductImageId: dataImage.ProductImageId,
				ImageUrl:       dataImage.ImageUrl,
			})
		}
	}

	resp = dataMap
	return
}

func (r salesRepository) CountListProduct(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error) {
	var (
		argIndex = 1
		args     = []interface{}{}
	)

	// Construct the SQL query
	query := `SELECT COUNT(DISTINCT sp.id)
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE 1 = 1`

	if filter.Search != "" {
		query += fmt.Sprintf(" AND sp.product_name ILIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.SalesId != "" {
		query += fmt.Sprintf(" AND sp.sales_id = $%d", argIndex)
		args = append(args, filter.SalesId)
		argIndex++
	}

	if filter.SubCategoryProduct != "" {
		query += fmt.Sprintf(" AND sp.sub_category_product = $%d", argIndex)
		args = append(args, filter.SubCategoryProduct)
		argIndex++
	}

	if filter.StatusProduct != "" {
		query += fmt.Sprintf(" AND statusProduct.status = $%d", argIndex)
		args = append(args, filter.StatusProduct)
		argIndex++
	}

	if filter.BestProduct != "" {
		bestProduct, errParse := strconv.ParseBool(filter.BestProduct)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.best_product = $%d", argIndex)
		args = append(args, bestProduct)
		argIndex++
	}

	if filter.IsActive != "" {
		isActive, errParse := strconv.ParseBool(filter.IsActive)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.is_active = $%d", argIndex)
		args = append(args, isActive)
		argIndex++
	}

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) CheckExistsProductId(ctx context.Context, id, salesId string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.sales_product WHERE id = $1 AND sales_id = $2)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, id, salesId).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) DetailSalesProduct(ctx context.Context, id string) (resp response.ProductDetailResp, err error) {
	var (
		dataImage    response.ProductImageDetail
		dataMapImage = make(map[string]*response.ProductImageDetail)
	)

	query := `SELECT sp.id, sp.sub_category_product, sp.product_name, sp.price, sp.tdp, sp.installment,
			sp.best_product, sp.is_active, spd.id, spd.description, statusProduct.status, spi.id, 
			spi.image_url, COALESCE(mp.id, '') AS province_id, COALESCE(mp.name, '') AS province_name, 
			COALESCE(mc.id, '') AS city_id, COALESCE(mc.name, '') AS city_name
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE sp.id = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, id)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&resp.IdProduct, &resp.ProductSubCategory,
			&resp.ProductName, &resp.Price, &resp.TDP, &resp.Installment, &resp.BestProduct,
			&resp.IsActive, &resp.IdDescription, &resp.Description, &resp.Status, &dataImage.ProductImageId,
			&dataImage.ImageUrl, &resp.ProvinceId, &resp.ProvinceName, &resp.CityId, &resp.CityName); err != nil {
			err = sharedError.HandleError(err)
		}

		if _, ok := dataMapImage[dataImage.ProductImageId]; !ok {
			if dataImage.ProductImageId != "" {
				resp.ProductImages = append(resp.ProductImages, response.ProductImageDetail{
					ProductImageId: dataImage.ProductImageId,
					ImageUrl:       dataImage.ImageUrl,
				})
			}

			dataMapImage[dataImage.ProductImageId] = &dataImage
		}
	}

	return
}

func (r salesRepository) GetCountDataImageByProductId(ctx context.Context, productId string) (count int, err error) {
	query := `
	SELECT
		COUNT(id)
	FROM 
		product_marketing.sales_product_image 
	WHERE product_id = $1`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, productId).Scan(&count); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) UpdateProductSales(ctx context.Context, tx *sql.Tx, param modelRequest.UpdateProductSales) (err error) {
	stmtProduct, err := tx.PrepareContext(ctx, `UPDATE product_marketing.sales_product SET product_name = $1, 
            price = $2, tdp = $3, installment = $4, best_product = $5, city_id = $6, slug = $7, sub_category_product = $8,
            updated_at = NOW() WHERE id = $9 `)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductDesc, err := tx.PrepareContext(ctx, `UPDATE product_marketing.sales_product_description 
			SET description = $1 WHERE id = $2`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	stmtProductImg, err := tx.PrepareContext(ctx, `UPDATE product_marketing.sales_product_image SET image_url = $1, is_active = $2 WHERE id = $3`)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProduct.ExecContext(ctx, param.ProductName, param.Price, param.TDP, param.Installment, param.BestProduct,
		param.CityId, param.Slug, param.ProductSubCategory, param.ProductId); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	if _, err = stmtProductDesc.ExecContext(ctx, param.Description, param.IdDescription); err != nil {
		err = sharedError.HandleError(err)
		return
	}

	for _, v := range param.Images {
		if _, err = stmtProductImg.ExecContext(ctx, v.ImageUrl, v.IsActive, v.ProductId); err != nil {
			err = sharedError.HandleError(err)
			return
		}
	}

	return
}

func (r salesRepository) ChangeStatusProductSales(ctx context.Context, tx *sql.Tx, param modelRequest.UpdateProductSales) (err error) {
	query := `INSERT INTO product_marketing.sales_product_status (id, product_id, status) VALUES ($1, $2, $3)`
	stmtProductStatus, err := tx.PrepareContext(ctx, query)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}

	logger.LogInfo(constant.QUERY, query)
	if _, err = stmtProductStatus.ExecContext(ctx, sharedRepo.GenerateUuidAsIdTable().String(), param.ProductId, param.Status); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) GetListProductSalesPublic(ctx context.Context, filter sharedModel.QueryRequest) (resp map[string]*response.ProductListSalesPublic, err error) {
	var (
		data      response.ProductListSalesPublic
		dataImage response.ProductImagePublic
		dataMap   = make(map[string]*response.ProductListSalesPublic)
		argIndex  = 1
		offset    = (filter.Page - 1) * filter.Limit
		args      = []interface{}{}
	)

	// Construct the SQL query
	query := `SELECT sp.id, COALESCE(sp.sub_category_product, ''), sp.product_name, sp.slug, sp.price, sp.tdp, sp.installment,
			sp.best_product, spd.description, statusProduct.status, spi.id, spi.image_url, COALESCE(mp.name, '') AS province_name, 
			COALESCE(mc.name, '') AS city_name, sp.created_at, sp.updated_at
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE 1 = 1 AND sp.is_active = TRUE`

	if filter.PublicAccess != "" {
		query += fmt.Sprintf(" AND sp.public_access = $%d", argIndex)
		args = append(args, filter.PublicAccess)
		argIndex++
	}

	if filter.Search != "" {
		query += fmt.Sprintf(" AND sp.product_name ILIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.SubCategoryProduct != "" {
		query += fmt.Sprintf(" AND sp.sub_category_product = $%d", argIndex)
		args = append(args, filter.SubCategoryProduct)
		argIndex++
	}

	if filter.StatusProduct != "" {
		query += fmt.Sprintf(" AND statusProduct.status = $%d", argIndex)
		args = append(args, filter.StatusProduct)
		argIndex++
	}

	if filter.BestProduct != "" {
		bestProduct, errParse := strconv.ParseBool(filter.BestProduct)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.best_product = $%d", argIndex)
		args = append(args, bestProduct)
		argIndex++
	}

	if filter.MinimumPrice != 0 && filter.MaximumPrice != 0 {
		query += fmt.Sprintf(" AND sp.price BETWEEN $%d AND $%d", argIndex, argIndex+1)
		args = append(args, filter.MinimumPrice, filter.MaximumPrice)
		argIndex += 2
	}

	if filter.SortBy != "" {
		query += " ORDER BY " + filter.SortBy

		if filter.SortOrder != "" {
			query += " " + filter.SortOrder
		}
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.Limit, offset)

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, args...)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&data.IdProduct, &data.ProductSubCategory, &data.ProductName, &data.Slug, &data.Price,
			&data.TDP, &data.Installment, &data.BestProduct, &data.Description, &data.Status, &dataImage.ProductImageId,
			&dataImage.ImageUrl, &data.ProvinceName, &data.CityName, &data.CreatedAt,
			&data.UpdatedAt); err != nil {
			err = sharedError.HandleError(err)
			return
		}

		if _, ok := dataMap[data.IdProduct]; !ok {
			dataMap[data.IdProduct] = &response.ProductListSalesPublic{
				ProductName:        data.ProductName,
				Slug:               data.Slug,
				Price:              data.Price,
				ProductSubCategory: data.ProductSubCategory,
				TDP:                data.TDP,
				Installment:        data.Installment,
				ProvinceName:       data.ProvinceName,
				CityName:           data.CityName,
				Description:        data.Description,
				Status:             data.Status,
				ProductImages:      []response.ProductImagePublic{},
			}
		}

		if dataImage.ProductImageId != "" {
			dataMap[data.IdProduct].ProductImages = append(dataMap[data.IdProduct].ProductImages, response.ProductImagePublic{
				ProductImageId: dataImage.ProductImageId,
				ImageUrl:       dataImage.ImageUrl,
			})
		}
	}

	resp = dataMap
	return
}

func (r salesRepository) CountListProductSalesPublic(ctx context.Context, filter sharedModel.QueryRequest) (count int, err error) {
	var (
		argIndex = 1
		args     = []interface{}{}
	)

	// Construct the SQL query
	query := `SELECT COUNT(DISTINCT sp.id)
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE 1 = 1 AND sp.is_active = TRUE`

	if filter.PublicAccess != "" {
		query += fmt.Sprintf(" AND sp.public_access = $%d", argIndex)
		args = append(args, filter.PublicAccess)
		argIndex++
	}

	if filter.Search != "" {
		query += fmt.Sprintf(" AND sp.product_name ILIKE $%d", argIndex)
		args = append(args, "%"+filter.Search+"%")
		argIndex++
	}

	if filter.SubCategoryProduct != "" {
		query += fmt.Sprintf(" AND sp.sub_category_product = $%d", argIndex)
		args = append(args, filter.SubCategoryProduct)
		argIndex++
	}

	if filter.StatusProduct != "" {
		query += fmt.Sprintf(" AND statusProduct.status = $%d", argIndex)
		args = append(args, filter.StatusProduct)
		argIndex++
	}

	if filter.BestProduct != "" {
		bestProduct, errParse := strconv.ParseBool(filter.BestProduct)
		if errParse != nil {
			err = sharedError.New(http.StatusBadRequest, errParse.Error(), errParse)
			return
		}

		query += fmt.Sprintf(" AND sp.best_product = $%d", argIndex)
		args = append(args, bestProduct)
		argIndex++
	}

	if filter.MinimumPrice != 0 && filter.MaximumPrice != 0 {
		query += fmt.Sprintf(" AND sp.price BETWEEN $%d AND $%d", argIndex, argIndex+1)
		args = append(args, filter.MinimumPrice, filter.MaximumPrice)
		argIndex += 2
	}

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, args...).Scan(&count); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}

func (r salesRepository) DetailSalesProductPublic(ctx context.Context, subdomain, slug string) (resp response.ProductDetailPublicResp, err error) {
	var (
		dataImage    response.ProductImageDetailPublic
		dataMapImage = make(map[string]*response.ProductImageDetailPublic)
	)

	query := `SELECT sp.id, sp.sub_category_product, sp.product_name, sp.price, sp.tdp, sp.installment,
			sp.best_product, spd.description, statusProduct.status, spi.id, 
			spi.image_url, COALESCE(mp.name, '') AS province_name, COALESCE(mc.name, '') AS city_name
			FROM product_marketing.sales_product sp
			LEFT JOIN product_marketing.sales_product_description spd ON sp.id = spd.product_id
			LEFT JOIN LATERAL (
				SELECT status
				FROM product_marketing.sales_product_status
				WHERE product_id = sp.id
				GROUP BY product_id, created_at, status
				ORDER BY created_at DESC LIMIT 1
			) statusProduct on true
			LEFT JOIN product_marketing.sales_product_image spi ON sp.id = spi.product_id
			LEFT JOIN product_marketing.master_city mc ON sp.city_id = mc.id
			LEFT JOIN product_marketing.master_province mp ON mc.province_id = mp.id
			WHERE sp.public_access = $1 AND sp.slug = $2 AND sp.is_active = TRUE`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain, slug)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&resp.IdProduct, &resp.ProductSubCategory, &resp.ProductName, &resp.Price,
			&resp.TDP, &resp.Installment, &resp.BestProduct, &resp.Description, &resp.Status,
			&dataImage.ProductImageId, &dataImage.ImageUrl, &resp.ProvinceName, &resp.CityName); err != nil {
			err = sharedError.HandleError(err)
		}

		if _, ok := dataMapImage[dataImage.ProductImageId]; !ok {
			if dataImage.ProductImageId != "" {
				resp.ProductImages = append(resp.ProductImages, response.ProductImageDetailPublic{
					ProductImageId: dataImage.ProductImageId,
					ImageUrl:       dataImage.ImageUrl,
				})
			}

			dataMapImage[dataImage.ProductImageId] = &dataImage
		}
	}

	return
}

func (r salesRepository) CheckExistsSlugProductSales(ctx context.Context, slug string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.sales_product WHERE slug = $1)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, slug).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
