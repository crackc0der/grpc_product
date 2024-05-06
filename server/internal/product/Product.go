package product

type Category struct {
	ID           int    `db:"id"`
	CategoryName string `db:"categoryName"`
}

type Product struct {
	ID              int      `db:"id"`
	ProductName     string   `db:"productName"`
	ProductCategory Category `db:"categoryName"`
	Price           float32  `db:"price"`
}
