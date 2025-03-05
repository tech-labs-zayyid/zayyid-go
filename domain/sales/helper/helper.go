package helper

func (p PageCategoryType) PageCategory() string {
	pageCategory := [...]string{
		"cars-sales",
		"property-sales",
		"bootcamp-sales",
	}

	return pageCategory[p]
}

func (s StatusProduct) StatusProduct() string {
	statusProductList := [...]string{
		"listed",
		"booked",
		"sold",
	}

	return statusProductList[s]
}

func (s StatusProductStr) IsValid() bool {
	status := map[StatusProductStr]bool{
		PRODUCT_LISTED_STRING: true,
		PRODUCT_BOOKED_STRING: true,
		PRODUCT_SOLD_STRING:   true,
	}

	return status[s]
}
