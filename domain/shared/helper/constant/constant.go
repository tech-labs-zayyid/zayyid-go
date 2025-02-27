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
	PAGE        = "page"
	LIMIT       = "limit"
	SORT_BY     = "sort_by"
	SORT_ORDER  = "sort_order"
	SEARCH      = "search"
	STATUS      = "status"
	PROVINCE_ID = "province_id"
)

const (
	ErrInvalidRequest       = "error invalid request"
	ErrMaximumUploadGallery = "you have reached the maximum limit for uploading gallery"
	ErrMaximumUploadBanner  = "you have reached the maximum limit for uploading banner"
	ErrRequestGallery       = "error invalid request gallery is empty"
	ErrRequestBanner        = "error invalid request banner is empty"
	ErrRequestProduct       = "error invalid request image product is empty"
	ErrDataNotFound         = "data not found"
)
