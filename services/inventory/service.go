package inventory

import (
	"context"
	"time"

	"github.com/nrednav/cuid2"
	"github.com/yash91989201/ecomm-monorepo/common/types"
)

type Service interface {
	InsertProduct(ctx context.Context, name, image, category, description string, rating, numReviews, countInStock int64, price float32) (*types.Product, error)
	SelectProductById(ctx context.Context, id string) (*types.Product, error)
	SelectAllProduct(ctx context.Context) ([]*types.Product, error)
	DeleteProductById(ctx context.Context, id string) error
}

type inventoryService struct {
	r Repository
}

func New(r Repository) Service {
	return &inventoryService{r: r}
}

func (s *inventoryService) InsertProduct(ctx context.Context, name, image, category, description string, rating, numReviews, countInStock int64, price float32) (*types.Product, error) {
	product := &types.Product{
		Id:           cuid2.Generate(),
		Name:         name,
		Image:        image,
		Category:     category,
		Description:  description,
		Rating:       rating,
		NumReviews:   numReviews,
		CountInStock: countInStock,
		Price:        price,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	p, err := s.r.InsertProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *inventoryService) SelectProductById(ctx context.Context, id string) (*types.Product, error) {
	return s.r.SelectProductById(ctx, id)
}

func (s *inventoryService) SelectAllProduct(ctx context.Context) ([]*types.Product, error) {
	return s.r.SelectAllProduct(ctx)
}

func (s *inventoryService) DeleteProductById(ctx context.Context, id string) error {
	return s.r.DeleteProductById(ctx, id)
}
