package category

import (
	"context"
	"fmt"
	product_grpc "server/api/note_v1"
)

type ServiceCategory struct {
	repository *RepositoryCategory
}

func NewServiceCategory(repository *RepositoryCategory) *ServiceCategory {
	return &ServiceCategory{repository: repository}
}

func (r *ServiceCategory) GetCategories(ctx context.Context) (*product_grpc.AllCategoryMessage, error) {
	categories, err := r.repository.SelectCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in service's method GetCategories: %w", err)
	}

	return categories, nil
}

func (r *ServiceCategory) AddCategory(ctx context.Context, category *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	category, err := r.repository.InsertCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("error in service's method AddCategory: %w", err)
	}

	return category, nil
}
