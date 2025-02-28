package helper

type PageCategoryType int
type StatusProduct int

const (
	CarsSalesProductCategoryPage PageCategoryType = iota
	PropertiesSalesProductCategoryPage
	BootcampSalesProductCategoryPage
)

const (
	ProductListed StatusProduct = iota
	ProductBooked
	ProductSold
)
