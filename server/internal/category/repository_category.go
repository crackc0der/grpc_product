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

func (r *RepositoryCategory) SelectCategories(_ context.Context) (*product_grpc.AllCategoryMessage, error) {
	query := "SELECT * FROM category"

	var categories []*product_grpc.CategoryMessage

	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method SelectCategories: %w", err)
	}

	allCategories := &product_grpc.AllCategoryMessage{
		Categories: categories,
	}
	return allCategories, nil
}

func (r *RepositoryCategory) InsertCategory(_ context.Context, category *product_grpc.CategoryMessage) (*product_grpc.CategoryMessage, error) {
	query := "INSERT INTO category (category_id, category_name) VALUES ($1, $2)"
	fmt.Println(category)
	_, err := r.db.Exec(query, category.Id, category.CategoryName)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method InsertCategory: %w", err)
	}

	return category, nil
}
