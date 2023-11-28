package service

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"github.com/khairulharu/marketplace/dto"
)

type cartService struct {
	cartRepository domain.CartRepository
}

func NewCart(cartRepository domain.CartRepository) domain.CartService {
	return &cartService{
		cartRepository: cartRepository,
	}
}

func (c cartService) AddProduct(ctx context.Context, req dto.CartReq) dto.Response {
	var cart = domain.Cart{
		UserID:    req.UserID,
		ProductID: req.ProductID,
	}
	if err := c.cartRepository.Insert(ctx, &cart); err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}

func (c cartService) GetCart(ctx context.Context, req dto.UserCartReq) dto.Response {
	var cartsRes []dto.Cart
	carts, err := c.cartRepository.FindByUserID(ctx, req.UserID)
	if err != nil {
		return dto.Response{
			Code:    "400",
			Massage: "INVALID",
			Error:   err.Error(),
		}
	}

	for _, v := range carts {
		cartsRes = append(cartsRes, dto.Cart{
			ID:        v.ID,
			UserID:    v.UserID,
			ProductID: v.ProductID,
		})
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
		Data:    cartsRes,
	}
}
