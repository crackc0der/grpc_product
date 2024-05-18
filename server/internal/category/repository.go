package category

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SelectCategories(_ context.Context) ([]Category, error) {
	query := `SELECT * FROM category`

	var categories []Category

	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method SelectCategories: %w", err)
	}

	return categories, nil
}

func (r *Repository) InsertCategory(_ context.Context, cat *Category) (*Category, error) {
	query := `INSERT INTO category (category_name) VALUES ($1) RETURNING category_id, category_name`

	var category Category

	err := r.db.QueryRowx(query, cat.CategoryName).StructScan(&category)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method InsertCategory: %w", err)
	}

	return &category, nil
}

func (r *Repository) UpdateCategory(_ context.Context, cat *Category) (*Category, error) {
	query := `UPDATE category SET category_name=$1 WHERE category_id=$2 RETURNING category_id, category_name`

	var category Category

	err := r.db.QueryRowx(query, cat.CategoryName, cat.CategoryID).StructScan(&category)
	if err != nil {
		return nil, fmt.Errorf("error in repository's method UpdateCategory: %w", err)
	}

	return &category, nil
}

func (r *Repository) DeleteCategory(_ context.Context, id int64) (bool, error) {
	query := `DELETE FROM category WHERE category_id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return false, fmt.Errorf("error in repository's method DeleteCategory: %w", err)
	}

	return true, nil
}
