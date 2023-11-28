package service

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"github.com/khairulharu/marketplace/dto"
)

type productService struct {
	productRepository domain.ProductRepository
}

func NewProduct(productRepository domain.ProductRepository) domain.ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p productService) ShowProduct(ctx context.Context) dto.Response {
	var resProducts []dto.Product
	products, err := p.productRepository.GetAll(ctx)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   err.Error(),
		}
	}

	for _, v := range products {
		resProducts = append(resProducts, dto.Product{
			ID:       v.ID,
			Name:     v.Name,
			Price:    v.Price,
			Category: v.Category,
		})
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
		Data:    resProducts,
	}

}
