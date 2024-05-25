package category

import (
	"context"
	"fmt"
)

type RepositoryInterface interface {
	SelectCategories(_ context.Context) ([]Category, error)
	InsertCategory(_ context.Context, cat *Category) (*Category, error)
	UpdateCategory(_ context.Context, cat *Category) (*Category, error)
	DeleteCategory(_ context.Context, id int64) (bool, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetCategories(ctx context.Context) ([]Category, error) {
	categories, err := s.repository.SelectCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in Server's service.GetCategories: %w", err)
	}

	return categories, nil
}

func (s *Service) AddCategory(ctx context.Context, category *Category) (*Category, error) {
	category, err := s.repository.InsertCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("error in Server's service.AddCategory: %w", err)
	}

	return category, nil
}

func (s *Service) UpdateCategory(ctx context.Context, category *Category) (*Category, error) {
	category, err := s.repository.UpdateCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("error in Server's service.UpdateCategory: %w", err)
	}

	return category, nil
}

func (s *Service) DeleteCategory(ctx context.Context, id int64) (bool, error) {
	result, err := s.repository.DeleteCategory(ctx, id)
	if err != nil {
		return result, fmt.Errorf("error in Server's service.DeleteCategory: %w", err)
	}

	return result, nil
}
