package repository

import (
	"context"
	"zayyid-go/domain/sales/model/response"
	"zayyid-go/domain/shared/helper/constant"
	sharedError "zayyid-go/domain/shared/helper/error"
	"zayyid-go/infrastructure/logger"
)

func (r salesRepository) HomeData(ctx context.Context, subdomain string) (resp response.DataHome, err error) {
	var (
		dataGallery        response.GalleryDataHomeResp
		dataBanner         response.BannerHomeResp
		dataSocialMedia    response.DataListSocialMediaHome
		dataTemplate       response.DataListTemplateHome
		dataTestimony      response.TestimonyListHome
		dataMapGallery     = make(map[string]*response.GalleryDataHomeResp)
		dataMapBanner      = make(map[string]*response.BannerHomeResp)
		dataMapSocialMedia = make(map[string]*response.DataListSocialMediaHome)
		dataMapTemplate    = make(map[string]*response.DataListTemplateHome)
		dataMapTestimony   = make(map[string]*response.TestimonyListHome)
	)

	query := `SELECT 
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
				product_marketing.sales_gallery sg
			LEFT JOIN LATERAL (
				SELECT id, image_url, description
				FROM product_marketing.sales_banner
				WHERE public_access = sg.public_access 
				AND is_active = TRUE
				LIMIT 4
			) sb ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, social_media_name, user_account, link_embed
				FROM product_marketing.sales_social_media
				WHERE public_access = sg.public_access 
				AND is_active = TRUE
				LIMIT 4
			) scm ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, color_plate_id
				FROM product_marketing.sales_template
				WHERE public_access = sg.public_access 
				AND is_active = TRUE
				LIMIT 2 
			) st ON TRUE
			LEFT JOIN LATERAL (
				SELECT id, fullname, description, photo_url
				FROM product_marketing.sales_testimony
				WHERE public_access = sg.public_access 
				AND is_active = TRUE
				LIMIT 6
			) sty ON TRUE 
			WHERE sg.public_access = $1`

	logger.LogInfo(constant.QUERY, query)
	rows, err := r.database.QueryContext(ctx, query, subdomain)
	if err != nil {
		err = sharedError.HandleError(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&dataGallery.IdGallery, &dataGallery.ImageUrl, &dataBanner.IdBanner, &dataBanner.ImageUrl,
			&dataBanner.Description, &dataSocialMedia.IdSocialMedia, &dataSocialMedia.SocialMediaName,
			&dataSocialMedia.UserAccount, &dataSocialMedia.LinkEmbed, &dataTemplate.IdTemplate, &dataTemplate.ColorPlate,
			&dataTestimony.IdTestimony, &dataTestimony.Name, &dataTestimony.Description, &dataTestimony.PhotoUrl)
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
	}

	return
}

func (r salesRepository) CheckExistsSubdomainSales(ctx context.Context, subdomain string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product_marketing.users WHERE subdomain = $1)`

	logger.LogInfo(constant.QUERY, query)
	if err = r.database.QueryRowContext(ctx, query, subdomain).Scan(&exists); err != nil {
		err = sharedError.HandleError(err)
	}

	return
}
