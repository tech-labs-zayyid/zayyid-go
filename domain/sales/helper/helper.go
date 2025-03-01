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
