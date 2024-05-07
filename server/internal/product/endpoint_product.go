package product

import (
	"context"
	"log/slog"

	product_grpc "server/api/note_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Endpoint struct {
	service *Service
	log     *slog.Logger
	product_grpc.UnimplementedProductServer
}

func NewEndpointProduct(service *Service, log *slog.Logger) *Endpoint {
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetProducts(ctx context.Context, _ *emptypb.Empty) (*product_grpc.AllProductMessage, error) {
	products, err := e.service.GetProducts(ctx)
	if err != nil {
		e.log.Error("error in Endpoint's method GetProducts: " + err.Error())

		return nil, err
	}

	return products, nil
}

func (e *Endpoint) GetProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductMessage, error) {
	product, err := e.service.GetProduct(ctx, id)
	if err != nil {
		e.log.Error("error in Endpoint's method GetProduct: " + err.Error())

		return nil, err
	}

	return product, nil
}

func (e *Endpoint) AddProduct(ctx context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	product, err := e.service.AddProduct(ctx, product)
	if err != nil {
		e.log.Error("error in Endpoint's method AddProduct: " + err.Error())

		return nil, err
	}

	return product, nil
}

func (e *Endpoint) DeleteProduct(ctx context.Context, id *product_grpc.ProductRequest) (*product_grpc.ProductResponse, error) {
	res, err := e.service.DeleteProduct(ctx, id)
	if err != nil {
		e.log.Error("error in Endpoint's method DeleteProduct: " + err.Error())

		return nil, err
	}

	return &product_grpc.ProductResponse{Deleted: res.GetDeleted()}, nil
}

func (e *Endpoint) UpdateProduct(ctx context.Context, product *product_grpc.ProductMessage) (*product_grpc.ProductMessage, error) {
	prod, err := e.service.UpdateProduct(ctx, product)
	if err != nil {
		e.log.Error("error in Endpoint's method UpdateProduct: " + err.Error())

		return nil, err
	}

	return prod, nil
}
