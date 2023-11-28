package repository

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"gorm.io/gorm"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) domain.CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (c cartRepository) FindByUserID(ctx context.Context, id int64) (carts []domain.Cart, err error) {
	err = c.db.Debug().WithContext(ctx).Table("carts").Where("user_id=?", id).Find(&carts).Error
	return
}

func (c cartRepository) Insert(ctx context.Context, cart *domain.Cart) error {
	err := c.db.Debug().WithContext(ctx).Table("carts").Save(cart).Error
	return err
}
