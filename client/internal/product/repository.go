package product

import (
	product_grpc "client/api/note_v1"
	"context"
	"fmt"
)

type Repository struct {
	client product_grpc.ProductClient
}

func NewRepository(cl product_grpc.ProductClient) *Repository {
	return &Repository{client: cl}
}

func (r *Repository) SelectProducts(ctx context.Context) ([]*Product, error) {
	productResult := []*Product{}

	products, err := r.client.GetProducts(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.SelectProducts: %w", err)
	}

	for _, product := range products.GetProducts() {
		prod := &Product{
			ID:                product.GetId(),
			ProductName:       product.GetProductName(),
			ProductCategoryID: product.GetProductCategoryID(),
			ProductPrice:      product.GetProductPrice(),
		}

		productResult = append(productResult, prod)
	}

	return productResult, nil
}

func (r *Repository) SelectProduct(ctx context.Context, id int64) (*Product, error) {
	productRequest := &product_grpc.ProductRequest{Id: id}

	product, err := r.client.GetProduct(ctx, productRequest)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.SelectProduct: %w", err)
	}

	productResult := &Product{
		ID:                product.GetId(),
		ProductName:       product.GetProductName(),
		ProductCategoryID: product.GetProductCategoryID(),
		ProductPrice:      product.GetProductPrice(),
	}

	return productResult, nil
}

func (r *Repository) InsertProduct(ctx context.Context, prod *Product) (*Product, error) {
	productMessage := &product_grpc.ProductMessage{
		Id:                prod.ID,
		ProductName:       prod.ProductName,
		ProductCategoryID: prod.ProductCategoryID,
		ProductPrice:      prod.ProductPrice,
	}

	product, err := r.client.AddProduct(ctx, productMessage)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.InsertProduct: %w", err)
	}

	productResult := &Product{
		ID:                product.GetId(),
		ProductName:       product.GetProductName(),
		ProductCategoryID: product.GetProductCategoryID(),
		ProductPrice:      product.GetProductPrice(),
	}

	return productResult, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	catID := &product_grpc.ProductRequest{
		Id: id,
	}

	deleted, err := r.client.DeleteProduct(ctx, catID)
	if err != nil {
		return deleted.GetDeleted(), fmt.Errorf("error in API's repository.DeleteProduct: %w", err)
	}

	return deleted.GetDeleted(), nil
}

func (r *Repository) UpdateProduct(ctx context.Context, prod *Product) (*Product, error) {
	productMessage := &product_grpc.ProductMessage{
		Id:                prod.ID,
		ProductName:       prod.ProductName,
		ProductCategoryID: prod.ProductCategoryID,
		ProductPrice:      prod.ProductPrice,
	}

	product, err := r.client.UpdateProduct(ctx, productMessage)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.UpdateProduct: %w", err)
	}

	productResult := &Product{
		ID:                productMessage.GetId(),
		ProductName:       product.GetProductName(),
		ProductCategoryID: product.GetProductCategoryID(),
		ProductPrice:      product.GetProductPrice(),
	}

	return productResult, nil
}
