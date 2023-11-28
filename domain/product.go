package domain

import (
	"context"

	"github.com/khairulharu/marketplace/dto"
)

type Product struct {
	ID       int64
	Name     string
	Price    string
	Category string
}

type ProductRepository interface {
	GetAll(ctx context.Context) ([]Product, error)
	FindById(ctx context.Context, id int64) (Product, error)
	FindByName(ctx context.Context, name string) (Product, error)
}

type ProductService interface {
	ShowProduct(ctx context.Context) dto.Response
}
