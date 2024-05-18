package product

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"

	product_grpc "server/api/note_v1"
)

type ServiceInterface interface {
	GetProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, id int64) (*Product, error)
	AddProduct(ctx context.Context, prod *Product) (*Product, error)
	DeleteProduct(ctx context.Context, productID int64) (bool, error)
	UpdateProduct(ctx context.Context, product *Product) (*Product, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
	product_grpc.UnimplementedProductServer
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	//nolint
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetProducts(ctx context.Context, _ *emptypb.Empty) (*product_grpc.AllProductMessage, error) {
	allProductMessage := []*product_grpc.ProductMessage{}

	products, err := e.service.GetProducts(ctx)
	if err != nil {
		e.log.Error("error in endpoint.GetProducts: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.GetProducts: %w", err)
	}

	for _, product := range products {
		prod := &product_grpc.ProductMessage{
			Id:                product.ID,
			ProductName:       product.ProductName,
			ProductCategoryID: product.ProductCategoryID,
			ProductPrice:      product.ProductPrice,
		}

		allProductMessage = append(allProductMessage, prod)
	}

	return &product_grpc.AllProductMessage{Products: allProductMessage}, nil
}

func (e *Endpoint) GetProduct(ctx context.Context, productID *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := e.service.GetProduct(ctx, productID.GetId())
	if err != nil {
		e.log.Error("error in endpoint.GetProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.GetProduct: %w", err)
	}

	productResult := &product_grpc.ProductMessage{
		Id:                product.ID,
		ProductName:       product.ProductName,
		ProductCategoryID: product.ProductCategoryID,
		ProductPrice:      product.ProductPrice,
	}

	return productResult, nil
}

func (e *Endpoint) AddProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	productStruct := &Product{
		ID:                prod.GetId(),
		ProductName:       prod.GetProductName(),
		ProductCategoryID: prod.GetProductCategoryID(),
		ProductPrice:      prod.GetProductPrice(),
	}

	product, err := e.service.AddProduct(ctx, productStruct)
	if err != nil {
		e.log.Error("error in endpoint.AddProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.AddProduct: %w", err)
	}

	productResult := &product_grpc.ProductMessage{
		Id:                product.ID,
		ProductName:       product.ProductName,
		ProductCategoryID: product.ProductCategoryID,
		ProductPrice:      product.ProductPrice,
	}

	return productResult, nil
}

func (e *Endpoint) DeleteProduct(ctx context.Context, prodID *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	res, err := e.service.DeleteProduct(ctx, prodID.GetId())
	if err != nil {
		e.log.Error("error in endpoint.DeleteProduct: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.DeleteProduct: %w", err)
	}

	return &product_grpc.ProductResponse{Deleted: res}, nil
}

func (e *Endpoint) UpdateProduct(ctx context.Context, prod *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	fmt.Println(prod)
	productStruct := &Product{
		ID:                prod.GetId(),
		ProductName:       prod.GetProductName(),
		ProductCategoryID: prod.GetProductCategoryID(),
		ProductPrice:      prod.GetProductPrice(),
	}

	product, err := e.service.UpdateProduct(ctx, productStruct)
	if err != nil {
		e.log.Error("error in endpoint.UpdateProduct111111111111: " + err.Error())

		return nil, fmt.Errorf("error in endpoint.UpdateProduct: %w", err)
	}

	productResult := &product_grpc.ProductMessage{
		Id:                product.ID,
		ProductName:       product.ProductName,
		ProductCategoryID: product.ProductCategoryID,
		ProductPrice:      product.ProductPrice,
	}

	return productResult, nil
}
