package product

type Product struct {
	ID                int64   `db:"product_id"`
	ProductName       string  `db:"product_name"`
	ProductCategoryID int64   `db:"product_category_id"`
	ProductPrice      float64 `db:"product_price"`
}
