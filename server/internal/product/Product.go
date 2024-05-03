package product

type Category struct {
	Id           int
	CategoryName string
}

type Product struct {
	Id              int
	ProductName     string
	ProductCategory int
	Price           float32
}
