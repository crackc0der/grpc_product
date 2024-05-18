package product

import (
	"context"
	"fmt"
)

type RepositoryInterface interface {
	SelectProducts(_ context.Context) ([]Product, error)
	SelectProductByID(_ context.Context, productID int64) (*Product, error)
	InsertProduct(_ context.Context, prod *Product) (*Product, error)
	DeleteProductByID(_ context.Context, productID int64) (bool, error)
	UpdateProduct(_ context.Context, product *Product) (*Product, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetProducts(ctx context.Context) ([]Product, error) {
	products, err := s.repository.SelectProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting products in service's method GetProducts: %w", err)
	}

	return products, nil
}

func (s *Service) GetProduct(ctx context.Context, productID int64) (*Product, error) {
	product, err := s.repository.SelectProductByID(ctx, productID)
	if err != nil {
		return nil, fmt.Errorf("error in service.GetProduct: %w", err)
	}

	return product, nil
}

func (s *Service) AddProduct(ctx context.Context, prod *Product) (*Product, error) {
	product, err := s.repository.InsertProduct(ctx, prod)
	if err != nil {
		return nil, fmt.Errorf("error in service.AddProduct: %w", err)
	}

	return product, nil
}

func (s *Service) DeleteProduct(ctx context.Context, productID int64) (bool, error) {
	result, err := s.repository.DeleteProductByID(ctx, productID)
	if err != nil {
		return result, fmt.Errorf("error in service.DeleteProduct: %w", err)
	}

	return result, nil
}

func (s *Service) UpdateProduct(ctx context.Context, product *Product) (*Product, error) {
	product, err := s.repository.UpdateProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("error in service.UpdateProduct22222222222222: %w", err)
	}

	return product, nil
}
