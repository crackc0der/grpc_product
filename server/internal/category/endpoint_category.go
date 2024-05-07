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
		return nil, fmt.Errorf("error in endpoint's method GetAllCategories: %w", err)
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
