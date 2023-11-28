package domain

import (
	"context"

	"github.com/khairulharu/marketplace/dto"
)

type Cart struct {
	ID        int64
	UserID    int64
	ProductID int64
}

type CartRepository interface {
	Insert(ctx context.Context, cart *Cart) error
	FindByUserID(ctx context.Context, id int64) ([]Cart, error)
}

type CartService interface {
	AddProduct(ctx context.Context, req dto.CartReq) dto.Response
	GetCart(ctx context.Context, req dto.UserCartReq) dto.Response
}
