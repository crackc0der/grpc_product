package product

import "server/internal/category"

type Product struct {
	ID              int               `db:"id"`
	ProductName     string            `db:"productName"`
	ProductCategory category.Category `db:"categoryName"`
	Price           float32           `db:"price"`
}
