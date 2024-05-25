package product

type Product struct {
	ID                int64   `json:"productId"`
	ProductName       string  `json:"productName"`
	ProductCategoryID int64   `json:"productCategoryId"`
	ProductPrice      float64 `json:"productPrice"`
}
