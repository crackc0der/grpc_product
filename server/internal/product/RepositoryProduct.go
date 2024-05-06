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
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Repository{db: conn}, nil
}

func (r *Repository) SelectAllProducts(_ context.Context) ([]*Product, error) {
	query := "SELECT * FROM product"

	var products []*Product

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("error selecting products in repository's method SelectAllProducts: %w", err)
	}

	return products, nil
}

func (r *Repository) SelectProductByID(_ context.Context, id int) (*Product, error) {
	query := "SELECT * FROM products WHERE id=:id"

	var product *Product

	res := strconv.Itoa(id)

	err := r.db.Get(product, query, res)
	if err != nil {
		return nil, fmt.Errorf("error selecting product in repository's method SelectProductById: %w", err)
	}

	return product, nil
}

func (r *Repository) InsertProduct(_ context.Context, prod *Product) (*Product, error) {
	query := `INSERT INTO product (ProductName, ProductCategory, Price) VALUES(:ProductName,:ProductCategory,:Price)`

	_, err := r.db.NamedExec(query, prod)
	if err != nil {
		return nil, fmt.Errorf("error inserting product in repository's method InsertProduct: %w", err)
	}

	return prod, nil
}

func (r *Repository) DeleteProductByID(_ context.Context, id int) error {
	query := "DELETE * FROM product WHERE id=:id"

	prodID := strconv.Itoa(id)

	_, err := r.db.Exec(query, prodID)
	if err != nil {
		return fmt.Errorf("error deleting product in repository's mothod DeleteProductById: %w", err)
	}

	return nil
}

func (r *Repository) UpdateProduct(_ context.Context, product *Product) (*Product, error) {
	query := `UPDATE product SET id=:Id ProductName=:ProductName ProductCategory=:ProductCategory Price=:Price 
			RETURNING id, ProductName, ProductCategory, Price`

	var updatedProduct *Product

	err := r.db.QueryRowx(query, product).StructScan(updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("error updating product in repository's method UpdateProduct: %w", err)
	}

	return updatedProduct, nil
}
