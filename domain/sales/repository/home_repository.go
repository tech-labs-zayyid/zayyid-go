package repository

import (
	"context"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) HomeData(ctx context.Context, subdomain string) (resp response.DataHome, product map[string]*response.BestProduct, err error) {
	var (
		dataGallery          response.GalleryDataHomeResp
		dataBanner           response.BannerHomeResp
		dataSocialMedia      response.DataListSocialMediaHome
		dataTemplate         response.DataListTemplateHome
		dataTestimony        response.TestimonyListHome
		dataProduct          response.BestProduct
		dataProductImage     response.BestProductImage
		dataMapGallery       = make(map[string]*response.GalleryDataHomeResp)
		dataMapBanner        = make(map[string]*response.BannerHomeResp)
		dataMapSocialMedia   = make(map[string]*response.DataListSocialMediaHome)
		dataMapTemplate      = make(map[string]*response.DataListTemplateHome)
		dataMapTestimony     = make(map[string]*response.TestimonyListHome)
		dataMapProduct       = make(map[string]*response.BestProduct)
		dataMapImagesProduct = make(map[string]*response.BestProductImage)
	)

	query := `SELECT 
    			COALESCE(u.name, '') AS name,
    			COALESCE(u.whatsapp_number, '') AS whatsapp_number,
    			COALESCE(u.email, '') AS email,
    			COALESCE(u.description, '') AS description,
    			COALESCE(u.image_url, '') AS image_url,
    			COALESCE(sp.id, '') AS product_id,
    			COALESCE(sp.product_name, '') AS product_name,
    			COALESCE(sp.price, 0) AS price,
    			COALESCE(sp.sub_category_product, '') AS sub_category_product,
    			COALESCE(sp.tdp, 0) AS tdp,
    			COALESCE(sp.installment, 0) AS installment,
    			COALESCE(mc.id, '') AS city_id,
    			COALESCE(mc.name, '') AS city_name,
    			COALESCE(mp.id, '') AS province_id,
    			COALESCE(mp.name, '') AS province_name,
    			COALESCE(sp.best_product, FALSE) AS best_product,
    			COALESCE(sp.slug, '') AS slug,
    			COALESCE(spi.id, '') AS product_image_id,
    			COALESCE(spi.image_url, '') AS product_image_url,
    			COALESCE(spd.description, '') AS product_description,
				COALESCE(sg.id, '') AS gallery_id, 
				COALESCE(sg.image_url, '') AS gallery_image,
				COALESCE(sb.id, '') AS banner_id, 
				COALESCE(sb.image_url, '') AS banner_image, 
				COALESCE(sb.description, '') AS banner_description,
				COALESCE(scm.id, '') AS social_media_id, 
				COALESCE(scm.social_media_name, ''), 
				COALESCE(scm.user_account, ''), 
				COALESCE(scm.link_embed, ''),
				COALESCE(st.id, '') AS template_id, 
				COALESCE(st.color_plate_id, ''),
				COALESCE(sty.id, '') AS testimony_id, 
				COALESCE(sty.fullname, ''), 
				COALESCE(sty.description, ''), 
				COALESCE(sty.photo_url, '')
			FROM
				product_marketing.users u
			LEFT JOIN LATERAL (
				SELECT id, product_name, price, sub_category_product, tdp, installment,
				city_id, image_url, best_product, slug
				FROM product_marketing.sales_product
				WHERE public_access = u.username 
				AND is_active = TRUE AND best_product = TRUE
				LIMIT 4
			) sp ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, image_url
				FROM product_marketing.sales_product_image
				WHERE product_id = sp.id
				AND is_active = TRUE
			) spi ON TRUE
			LEFT JOIN LATERAL (
				SELECT description
				FROM product_marketing.sales_product_description
				WHERE product_id = sp.id
			) spd ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, name, province_id
				FROM product_marketing.master_city
				WHERE id = sp.city_id 
				AND is_active = TRUE
			) mc ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, name
				FROM product_marketing.master_province
				WHERE id = mc.province_id 
				AND is_active = TRUE
			) mp ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, image_url
				FROM product_marketing.sales_gallery
				WHERE public_access = u.username 
				AND is_active = TRUE
				LIMIT 4
			) sg ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, image_url, description
				FROM product_marketing.sales_banner
				WHERE public_access = u.username 
				AND is_active = TRUE
				LIMIT 4
			) sb ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, social_media_name, user_account, link_embed
				FROM product_marketing.sales_social_media
				WHERE public_access = u.username 
				AND is_active = TRUE
				LIMIT 4
			) scm ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, color_plate_id
				FROM product_marketing.sales_template
				WHERE public_access = u.username 
				AND is_active = TRUE
				LIMIT 1 
			) st ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, fullname, description, photo_url
				FROM product_marketing.sales_testimony
				WHERE public_access = u.username 
				AND is_active = TRUE
				LIMIT 6
			) sty ON TRUE 
			WHERE u.username = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&resp.Fullname, &resp.PhoneNumber, &resp.Email, &resp.Desc, &resp.UrlImage, &dataProduct.IdProduct, &dataProduct.ProductName, &dataProduct.Price, &dataProduct.ProductSubCategory,
			&dataProduct.TDP, &dataProduct.Installment, &dataProduct.CityId, &dataProduct.CityName, &dataProduct.ProvinceId,
			&dataProduct.ProvinceName, &dataProduct.BestProduct, &dataProduct.Slug, &dataProductImage.ProductImageId,
			&dataProductImage.ImageUrl, &dataProduct.Description, &dataGallery.IdGallery, &dataGallery.ImageUrl,
			&dataBanner.IdBanner, &dataBanner.ImageUrl, &dataBanner.Description, &dataSocialMedia.IdSocialMedia,
			&dataSocialMedia.SocialMediaName, &dataSocialMedia.UserAccount, &dataSocialMedia.LinkEmbed, &dataTemplate.IdTemplate,
			&dataTemplate.ColorPlate, &dataTestimony.IdTestimony, &dataTestimony.Name, &dataTestimony.Description, &dataTestimony.PhotoUrl)
		if err != nil {
			err = sharedError.HandleError(err)
			return
		}

		if _, ok := dataMapGallery[dataGallery.IdGallery]; !ok {
			if dataGallery.IdGallery != "" {
				resp.Gallery = append(resp.Gallery, response.GalleryDataHomeResp{
					IdGallery: dataGallery.IdGallery,
					ImageUrl:  dataGallery.ImageUrl,
				})
			}

			dataMapGallery[dataGallery.IdGallery] = &dataGallery
		}

		if _, ok := dataMapBanner[dataBanner.IdBanner]; !ok {
			if dataBanner.IdBanner != "" {
				resp.Banner = append(resp.Banner, response.BannerHomeResp{
					IdBanner:    dataBanner.IdBanner,
					ImageUrl:    dataBanner.ImageUrl,
					Description: dataBanner.Description,
				})
			}

			dataMapBanner[dataBanner.IdBanner] = &dataBanner
		}

		if _, ok := dataMapSocialMedia[dataSocialMedia.IdSocialMedia]; !ok {
			if dataSocialMedia.IdSocialMedia != "" {
				resp.SocialMedia = append(resp.SocialMedia, response.DataListSocialMediaHome{
					IdSocialMedia:   dataSocialMedia.IdSocialMedia,
					SocialMediaName: dataSocialMedia.SocialMediaName,
					UserAccount:     dataSocialMedia.UserAccount,
					LinkEmbed:       dataSocialMedia.LinkEmbed,
				})
			}

			dataMapSocialMedia[dataSocialMedia.IdSocialMedia] = &dataSocialMedia
		}

		if _, ok := dataMapTemplate[dataTemplate.IdTemplate]; !ok {
			if dataTemplate.IdTemplate != "" {
				resp.Template = append(resp.Template, response.DataListTemplateHome{
					IdTemplate: dataTemplate.IdTemplate,
					ColorPlate: dataTemplate.ColorPlate,
				})
			}

			dataMapTemplate[dataTemplate.IdTemplate] = &dataTemplate
		}

		if _, ok := dataMapTestimony[dataTestimony.IdTestimony]; !ok {
			if dataTestimony.IdTestimony != "" {
				resp.Testimony = append(resp.Testimony, response.TestimonyListHome{
					IdTestimony: dataTestimony.IdTestimony,
					Name:        dataTestimony.Name,
					Description: dataTestimony.Description,
					PhotoUrl:    dataTestimony.PhotoUrl,
				})
			}

			dataMapTestimony[dataTestimony.IdTestimony] = &dataTestimony
		}

		if _, ok := dataMapProduct[dataProduct.IdProduct]; !ok {
			dataMapProduct[dataProduct.IdProduct] = &response.BestProduct{
				ProductName:        dataProduct.ProductName,
				Price:              dataProduct.Price,
				ProductSubCategory: dataProduct.ProductSubCategory,
				TDP:                dataProduct.TDP,
				Installment:        dataProduct.Installment,
				ProvinceId:         dataProduct.ProvinceId,
				ProvinceName:       dataProduct.ProvinceName,
				CityId:             dataProduct.CityId,
				CityName:           dataProduct.CityName,
				BestProduct:        dataProduct.BestProduct,
				Slug:               dataProduct.Slug,
				Description:        dataProduct.Description,
				Images:             []response.BestProductImage{},
			}
		}

		if _, ok := dataMapImagesProduct[dataProductImage.ProductImageId]; !ok {
			if dataProductImage.ProductImageId != "" {
				dataMapProduct[dataProduct.IdProduct].Images = append(dataMapProduct[dataProduct.IdProduct].Images, response.BestProductImage{
					ProductImageId: dataProductImage.ProductImageId,
					ImageUrl:       dataProductImage.ImageUrl,
				})
			}

			dataMapImagesProduct[dataProductImage.ProductImageId] = &dataProductImage
		}
	}

	product = dataMapProduct
	return
}
