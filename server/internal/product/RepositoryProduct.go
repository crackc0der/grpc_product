package product

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db string, dsn string) (*Repository, error) {
	conn, err := sqlx.Connect(db, dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return &Repository{db: conn}, nil
}

func (r *Repository) SelectAllProducts(ctx context.Context) ([]*Product, error) {
	query := "SELECT * FROM product"
	products := []*Product{}

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("error selecting products in repository's method SelectAllProducts: %v", err)
	}

	return products, nil
}

func (r *Repository) SelectProductById(ctx context.Context, id int) (*Product, error) {
	query := "SELECT * FROM products WHERE id=:id"
	product := &Product{}

	res := strconv.Itoa(id)

	err := r.db.Get(product, query, res)
	if err != nil {
		return nil, fmt.Errorf("error selecting product in repository's method SelectProductById: %v", err)
	}

	return product, nil
}

func (r *Repository) InsertProduct(ctx context.Context, prod *Product) (*Product, error) {
	query := `INSERT INTO product (ProductName, ProductCategory, Price) VALUES(:ProductName,:ProductCategory,:Price)`
	product := &Product{}

	_, err := r.db.NamedExec(query, product)
	if err != nil {
		return nil, fmt.Errorf("error inserting product in repository's method InsertProduct: %v", err)
	}

	return product, nil
}

func (r *Repository) DeleteProductById(ctx context.Context, id int) error {
	query := "DELETE * FROM product WHERE id=:id"

	prodId := strconv.Itoa(id)

	_, err := r.db.Exec(query, prodId)

	if err != nil {
		return fmt.Errorf("error deleting product in repository's mothod DeleteProductById: %v", err)
	}

	return nil
}

func (r *Repository) UpdateProduct(ctx context.Context, product *Product) (*Product, error) {
	query := `UPDATE product SET id=:Id ProductName=:ProductName ProductCategory=:ProductCategory Price=:Price 
			RETURNING id, ProductName, ProductCategory, Price`
	updatedProduct := &Product{}

	err := r.db.QueryRowx(query, product).StructScan(updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("error updating product in repository's method UpdateProduct: %v", err)
	}

	return updatedProduct, nil
}
