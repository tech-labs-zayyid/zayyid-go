package constant

const (
	HEADER_API_KEY   = "X-Api-Key"
	HEADER_TIMESTAMP = "X-Timestamp"
	HEADER_SIGNATURE = "X-Signature"

	HEADER_INTEGRATION_TIME = "X-Integration-Time"

	PARAM_BASE_PATH = "base_path"
	PARAM_ENDPOINT  = "*"
)

const (
	SUCCESS          = "success"
	FAILED           = "failed"
	ERROR            = "error"
	ERRJWTVALIDATION = "no Keyfunc was provided."
)

type ContextKey string

const (
	FiberContext  ContextKey = "fiberCtx"
	HeaderContext ContextKey = "headerCtx"
)

const (
	RESPONSE = "response"
	QUERY    = "query"
	AUTH     = "auth"
)

const (
	DefaultTimeout = 10 // detik
)

const (
	PAGE                 = "page"
	LIMIT                = "limit"
	SORT_BY              = "sort_by"
	SORT_ORDER           = "sort_order"
	SEARCH               = "search"
	STATUS               = "status"
	PROVINCE_ID          = "province_id"
	IS_ACTIVE            = "is_active"
	SUB_CATEGORY_PRODUCT = "sub_category_product"
	BEST_PRODUCT         = "best_product"
	STATUS_PRODUCT       = "status_product"
	MINIMUM_PRICE        = "minimum_price"
	MAXIMUM_PRICE        = "maximum_price"
)

const (
	ErrInvalidRequest         = "error invalid request"
	ErrMaximumUploadGallery   = "you have reached the maximum limit for uploading gallery"
	ErrMaximumUploadBanner    = "you have reached the maximum limit for uploading banner"
	ErrRequestGallery         = "error invalid request gallery is empty"
	ErrRequestBanner          = "error invalid request banner is empty"
	ErrRequestProductImages   = "error invalid request image product not valid length"
	ErrDataNotFound           = "data not found"
	ErrInvalidSignatureKey    = "invalid signature key"
	ErrDataUserIdNotFound     = "error not found user id"
	ErrIdProductNotFound      = "id product not found"
	ErrDuplicateProductName   = "product name duplicate entry"
	ErrStatusInvalid          = "invalid status product"
	ErrStatusNotSettelment    = "status not settelment"
	ErrSlugNotFound           = "slug not found"
	ErrReferralCode           = "referral code not found"
	ErrIdSocialNotFound       = "social media not found"
	ErrIdGalleryNotFound      = "gallery not found"
	ErrIdBannerNotFound       = "banner not found"
	ErrIdTemplateNotFound     = "template not found"
	ErrLinkEmbedNotPermission = "link embed is not permission"
)

const (
	SETTLEMENT = "settlement"
	ACCEPT     = "accept"
)
