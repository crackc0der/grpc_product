package product

import (
	"context"
	"fmt"
	"strconv"

	_ "github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"

	product_grpc "server/api/note_v1"
)

type RepositoryProduct struct {
	db *sqlx.DB
}

func NewRepositoryProduct(db *sqlx.DB) *RepositoryProduct {
	return &RepositoryProduct{db: db}
}

func (r *RepositoryProduct) SelectProducts(_ context.Context) (*product_grpc.AllProductMessage, error) {
	query := "SELECT * FROM product"

	var products []*product_grpc.ProductMessage

	err := r.db.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("error selecting product in repository's method SelectAllProducts: %w", err)
	}

	allProducts := &product_grpc.AllProductMessage{
		Products: products,
	}

	return allProducts, nil
}

func (r *RepositoryProduct) SelectProductByID(_ context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	query := "SELECT * FROM product WHERE product_id=$1"

	product := &product_grpc.ProductMessage{}
	fmt.Println(id.GetId())
	err := r.db.Get(product, query, id.GetId())
	if err != nil {
		return nil, fmt.Errorf("error selecting product in repository's method SelectProductById: %w", err)
	}

	return product, nil
}

func (r *RepositoryProduct) InsertProduct(_ context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	query := `INSERT INTO product (product_name, product_category_id, product_price) VALUES($1, $2, $3) 
		RETURNING product_id, product_name, product_category_id, product_price`

	product := &product_grpc.ProductMessage{}

	err := r.db.QueryRowx(query, prod.ProductName, prod.ProductCategoryID, prod.ProductPrice).StructScan(product)
	if err != nil {
		return nil, fmt.Errorf("error inserting product in repository's method InsertProduct: %w", err)
	}

	return product, nil
}

func (r *RepositoryProduct) DeleteProductByID(_ context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	query := "DELETE FROM product WHERE product_id=$1"

	prodID := strconv.FormatInt(productID.GetId(), 10)

	_, err := r.db.Exec(query, prodID)
	if err != nil {
		return nil, fmt.Errorf("error deleting product in repository's mothod DeleteProductById: %w", err)
	}

	return &product_grpc.ProductResponse{Deleted: true}, nil
}

func (r *RepositoryProduct) UpdateProduct(_ context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	query := `UPDATE product SET product_name=$1, product_category_id=$2, product_price=$3 WHERE product_id=$4
			RETURNING product_id, product_name, product_category_id, product_price`

	updatedProduct := &product_grpc.ProductMessage{}

	err := r.db.QueryRowx(query, product.ProductName, product.ProductCategoryID, product.ProductPrice, product.Id).StructScan(updatedProduct)
	if err != nil {
		return nil, fmt.Errorf("error updating product in repository's method UpdateProduct: %w", err)
	}

	return updatedProduct, nil
}
