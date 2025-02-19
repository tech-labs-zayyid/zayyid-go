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
	DefaultTimeout = 5 // detik
)

const (
	PAGE       = "page"
	LIMIT      = "limit"
	SORT_BY    = "sort_by"
	SORT_ORDER = "sort_order"
	SEARCH     = "search"
	STATUS     = "status"
)
