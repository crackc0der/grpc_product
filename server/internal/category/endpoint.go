package category

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/protobuf/types/known/emptypb"

	product_grpc "server/api/note_v1"
)

type ServiceInterface interface {
	GetCategories(ctx context.Context) ([]Category, error)
	AddCategory(ctx context.Context, cat *Category) (*Category, error)
	UpdateCategory(ctx context.Context, cat *Category) (*Category, error)
	DeleteCategory(ctx context.Context, id int64) (bool, error)
}

type Endpoint struct {
	service ServiceInterface
	log     *slog.Logger
	product_grpc.UnimplementedCategoryServer
}

func NewEndpoint(service *Service, log *slog.Logger) *Endpoint {
	//nolint
	return &Endpoint{service: service, log: log}
}

func (e *Endpoint) GetCategories(ctx context.Context, _ *emptypb.Empty) (*product_grpc.AllCategoryMessage, error) {
	categoryMessage := []*product_grpc.CategoryMessage{}

	categories, err := e.service.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in endpoint's method GetCategories: %w", err)
	}

	for _, category := range categories {
		cat := &product_grpc.CategoryMessage{
			Id:           int64(category.CategoryID),
			CategoryName: category.CategoryName,
		}

		categoryMessage = append(categoryMessage, cat)
	}

	return &product_grpc.AllCategoryMessage{
		Categories: categoryMessage,
	}, nil
}

func (e *Endpoint) AddCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	categoryStruct := &Category{
		CategoryID:   int(cat.GetId()),
		CategoryName: cat.GetCategoryName(),
	}

	category, err := e.service.AddCategory(ctx, categoryStruct)
	if err != nil {
		return nil, fmt.Errorf("error in endpoint's method AddCategory: %w", err)
	}

	categoryMessage := &product_grpc.CategoryMessage{
		Id:           int64(category.CategoryID),
		CategoryName: category.CategoryName,
	}

	return categoryMessage, nil
}

func (e *Endpoint) UpdateCategory(ctx context.Context, cat *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	categoryStruct := &Category{
		CategoryID:   int(cat.GetId()),
		CategoryName: cat.GetCategoryName(),
	}

	category, err := e.service.UpdateCategory(ctx, categoryStruct)
	if err != nil {
		e.log.Error("error in Endpoint's method UpdateCategory: " + err.Error())

		return nil, fmt.Errorf("error in service's method UpdateCategory: %w", err)
	}

	categoryMessage := &product_grpc.CategoryMessage{
		Id:           int64(category.CategoryID),
		CategoryName: category.CategoryName,
	}

	return categoryMessage, nil
}

func (e *Endpoint) DeleteCategory(ctx context.Context, catID *product_grpc.CategoryRequest) (*product_grpc.CategoryResponse, error) {
	result, err := e.service.DeleteCategory(ctx, catID.GetId())
	if err != nil {
		e.log.Error("error in Endpoint's method DeleteCategory: " + err.Error())

		return nil, fmt.Errorf("error in service's method DeleteCategory: %w", err)
	}

	return &product_grpc.CategoryResponse{Deleted: result}, nil
}
