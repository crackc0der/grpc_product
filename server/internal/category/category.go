package category

type Category struct {
	ID           int    `db:"id"`
	CategoryName string `db:"categoryName"`
}
