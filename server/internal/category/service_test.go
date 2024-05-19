package category

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCategories(t *testing.T) {
	t.Parallel()

	type args struct {
		ctx context.Context
	}

	repo := new(MockRepositoryInterface)

	set := func(list []Category, err error) {
		repo.On("SelectCategories", mock.Anything).Return(list, err)
	}

	categories := []struct {
		args  args
		name  string
		want  []Category
		err   error
		setUp func()
	}{
		{
			args: args{ctx: context.Background()},
			name: "category",
			want: []Category{
				{
					CategoryID:   1,
					CategoryName: "category 1",
				},
				{
					CategoryID:   2,
					CategoryName: "category 2",
				},
			},
			err: nil,
			setUp: func() {
				set([]Category{
					{
						CategoryID:   1,
						CategoryName: "category 1",
					},
					{
						CategoryID:   2,
						CategoryName: "category 2",
					},
				}, nil)
			},
		},
	}

	for _, category := range categories {
		t.Run(category.name, func(t *testing.T) {
			t.Parallel()
			category.setUp()

			service := NewService(repo)

			categories, err := service.GetCategories(category.args.ctx)
			if !errors.Is(err, category.err) {
				t.Error(err)
			}

			assert.Equal(t, category.want, categories)
		})
	}
}

// func TestAddCategory() {

// }

// func TestUpdateCategory() {

// }

// func TestDeleteCategory() {

// }
