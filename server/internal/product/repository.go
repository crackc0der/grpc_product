package product

import (
	"context"
	"fmt"

	// Importing pgx for indirect use via another package.
	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectProducts(_ context.Context) ([]Product, error) {
	query := "SELECT * FROM product"

	var products []Product

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("error in Server's repository.SelectProducts: %w", err)
	}

	return products, nil
}

func (r *Repository) SelectProductByID(_ context.Context, productID int64) (*Product, error) {
	var product Product

	query := "SELECT * FROM product WHERE product_id=$1"

	err := r.db.Get(&product, query, productID)
	if err != nil {
		return nil, fmt.Errorf("error in Server's repository.SelectProductById: %w", err)
	}

	return &product, nil
}

func (r *Repository) InsertProduct(_ context.Context, prod *Product) (*Product, error) {
	query := `INSERT INTO product (product_name, product_category_id, product_price) VALUES($1, $2, $3) 
		RETURNING product_id, product_name, product_category_id, product_price`

	var product Product

	err := r.db.QueryRowx(query, prod.ProductName, prod.ProductCategoryID, prod.ProductPrice).StructScan(&product)
	if err != nil {
		return nil, fmt.Errorf("error in Server's repository.InsertProduct: %w", err)
	}

	return &product, nil
}

func (r *Repository) DeleteProductByID(_ context.Context, productID int64) (bool, error) {
	query := "DELETE FROM product WHERE product_id=$1"

	_, err := r.db.Exec(query, productID)
	if err != nil {
		return false, fmt.Errorf("error in Server's repository.DeleteProductById: %w", err)
	}

	return true, nil
}

func (r *Repository) UpdateProduct(_ context.Context, product *Product) (*Product, error) {
	query := `UPDATE product SET product_name=$1, product_category_id=$2, product_price=$3 WHERE product_id=$4
			RETURNING product_id, product_name, product_category_id, product_price`

	var updatedProduct Product

	err := r.db.QueryRowx(query, product.ProductName, product.ProductCategoryID, product.ProductPrice, product.ID).StructScan(&updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("error in Server's repository.UpdateProduct: %w", err)
	}

	return &updatedProduct, nil
}
