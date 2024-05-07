package category

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	product_grpc "server/api/note_v1"
)

type RepositoryCategory struct {
	db *sqlx.DB
}

func NewRepositoryCategory(db *sqlx.DB) *RepositoryCategory {
	return &RepositoryCategory{db: db}
}

func (r *RepositoryCategory) SelectAllCategories(_ context.Context) (*product_grpc.AllCategoryMessage, error) {
	query := "SELECT * FROM category"

	var categories *product_grpc.AllCategoryMessage

	err := r.db.Select(categories.GetCategories(), query)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method SelectAllCategoris: %w", err)
	}

	return categories, nil
}

func (r *RepositoryCategory) InsertCategory(_ context.Context, category *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	query := "INSERT INTO category (category_name) VALUE (:CategoryName)"

	_, err := r.db.NamedExec(query, category)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method InsertCategory: %w", err)
	}

	return category, nil
}
