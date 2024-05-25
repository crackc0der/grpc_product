package category

import (
	"context"
	"fmt"
)

type RepositoryInterface interface {
	SelectCategories(ctx context.Context) ([]*Category, error)
	InsertCategory(ctx context.Context, cat *Category) (*Category, error)
	UpdateCategory(ctx context.Context, cat *Category) (*Category, error)
	DeleteCategory(ctx context.Context, id int64) (bool, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetCategories(ctx context.Context) ([]*Category, error) {
	categories, err := s.repository.SelectCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.GetCategories: %w", err)
	}

	return categories, nil
}

func (s *Service) AddCategory(ctx context.Context, cat *Category) (*Category, error) {
	category, err := s.repository.InsertCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.AddCategory: %w", err)
	}

	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, cat *Category) (*Category, error) {
	category, err := s.repository.UpdateCategory(ctx, cat)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.UpdateCategory: %w", err)
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id int64) (bool, error) {
	result, err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return result, fmt.Errorf("error in API's service.DeleteCategory: %w", err)
	}

	return result, nil
}
