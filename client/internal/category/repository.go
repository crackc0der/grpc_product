package category

import (
	product_grpc "client/api/note_v1"
	"context"
	"fmt"
)

type Repository struct {
	client product_grpc.CategoryClient
}

func NewRepository(cl product_grpc.CategoryClient) *Repository {
	return &Repository{client: cl}
}

func (r *Repository) SelectCategories(ctx context.Context) ([]*Category, error) {
	categoriesResult := []*Category{}

	categories, err := r.client.GetCategories(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.SelectCategories: %w", err)
	}

	for _, category := range categories.GetCategories() {
		cat := &Category{
			CategoryID:   category.GetId(),
			CategoryName: category.GetCategoryName(),
		}

		categoriesResult = append(categoriesResult, cat)
	}

	return categoriesResult, nil
}

func (r *Repository) InsertCategory(ctx context.Context, cat *Category) (*Category, error) {
	categoryMessage := &product_grpc.CategoryMessage{
		Id:           cat.CategoryID,
		CategoryName: cat.CategoryName,
	}

	category, err := r.client.AddCategory(ctx, categoryMessage)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.InsertCategory: %w", err)
	}

	categoryResult := &Category{
		CategoryID:   category.GetId(),
		CategoryName: category.GetCategoryName(),
	}

	return categoryResult, nil
}

func (r *Repository) UpdateCategory(ctx context.Context, cat *Category) (*Category, error) {
	categoryMessage := &product_grpc.CategoryMessage{
		Id:           cat.CategoryID,
		CategoryName: cat.CategoryName,
	}

	category, err := r.client.UpdateCategory(ctx, categoryMessage)
	if err != nil {
		return nil, fmt.Errorf("error in API's repository.UpdateCategory: %w", err)
	}

	categoryResult := &Category{
		CategoryID:   category.GetId(),
		CategoryName: category.GetCategoryName(),
	}

	return categoryResult, nil
}

func (r *Repository) DeleteCategory(ctx context.Context, id int64) (bool, error) {
	catID := &product_grpc.CategoryRequest{
		Id: id,
	}

	result, err := r.client.DeleteCategory(ctx, catID)
	if err != nil {
		return result.GetDeleted(), fmt.Errorf("error in API's repository.DeleteCategory: %w", err)
	}

	return result.GetDeleted(), nil
}
