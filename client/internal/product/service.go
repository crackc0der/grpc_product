package product

import (
	"context"
	"fmt"
)

type RepositoryInterface interface {
	SelectProducts(ctx context.Context) ([]*Product, error)
	SelectProduct(ctx context.Context, id int64) (*Product, error)
	InsertProduct(ctx context.Context, prod *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id int64) (bool, error)
	UpdateProduct(ctx context.Context, prod *Product) (*Product, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetProducts(ctx context.Context) ([]*Product, error) {
	products, err := s.repository.SelectProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.GetProducts: %w", err)
	}

	return products, nil
}

func (s *Service) GetProduct(ctx context.Context, id int64) (*Product, error) {
	product, err := s.repository.SelectProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.GetProduct: %w", err)
	}

	return product, nil
}

func (s *Service) AddProduct(ctx context.Context, prod *Product) (*Product, error) {
	product, err := s.repository.InsertProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.AddProduct: %w", err)
	}

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, id int64) (bool, error) {
	result, err := s.repository.DeleteProduct(ctx, id)
	if err != nil {
		return result, fmt.Errorf("error in API's service.DeleteProduct: %w", err)
	}

	return result, nil
}

func (s *Service) UpdateProduct(ctx context.Context, prod *Product) (*Product, error) {
	product, err := s.repository.UpdateProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in API's service.UpdateProduct: %w", err)
	}

	return product, nil
}
