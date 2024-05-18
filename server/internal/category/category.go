package category

type Category struct {
	CategoryID   int    `db:"category_id"`
	CategoryName string `db:"category_name"`
}
