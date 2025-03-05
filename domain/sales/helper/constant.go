package helper

type PageCategoryType int
type StatusProduct int
type StatusProductStr string

const (
	CARS_SALES_PRODUCT_CATEGORY_PAGE PageCategoryType = iota
	PROPERTIES_SALES_PRODUCT_CATEGORY_PAGE
	BOOTCAMP_SALES_PRODUCT_CATEGORY_PAGE
)

const (
	PRODUCT_LISTED StatusProduct = iota
	PRODUCT_BOOKED
	PRODUCT_SOLD
)

const (
	PRODUCT_LISTED_STRING = "listed"
	PRODUCT_BOOKED_STRING = "booked"
	PRODUCT_SOLD_STRING   = "sold"
)
