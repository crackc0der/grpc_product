package category

import (
	"context"
	"fmt"
	"log/slog"

	product_grpc "server/api/note_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type EndpointCategory struct {
	service *ServiceCategory
	log     *slog.Logger
	product_grpc.UnimplementedCategoryServer
}

func NewEndpointCategory(service *ServiceCategory, log *slog.Logger) *EndpointCategory {
	return &EndpointCategory{service: service, log: log}
}

func (e *EndpointCategory) GetCategories(ctx context.Context, _ *emptypb.Empty) (*product_grpc.AllCategoryMessage, error) {
	categories, err := e.service.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in endpoint's method GetCategories: %w", err)
	}
	return categories, nil
}

func (e *EndpointCategory) AddCategory(ctx context.Context, category *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := e.service.AddCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("error in endpoint's method AddCategory: %w", err)
	}

	return category, nil
}

func (e *EndpointCategory) UpdateCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := e.service.UpdateCategory(ctx, cat)
	if err != nil {
		e.log.Error("error in Endpoint's method UpdateCategory: " + err.Error())

		return nil, fmt.Errorf("error in service's method UpdateCategory: %w", err)
	}

	return category, nil
}

func (e *EndpointCategory) DeleteCategory(ctx context.Context, id *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error) {
	result, err := e.service.DeleteCategory(ctx, id)
	if err != nil {
		e.log.Error("error in Endpoint's method DeleteCategory: " + err.Error())

		return nil, fmt.Errorf("error in service's method DeleteCategory: %w", err)
	}

	return result, nil
}
