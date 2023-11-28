package repository

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p productRepository) FindById(ctx context.Context, id int64) (product domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("products").Where("id=?", id).First(&product).Error
	return
}

func (p productRepository) FindByName(ctx context.Context, name string) (product domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("products").Where("name=?", name).First(&product).Error
	return
}

func (p productRepository) GetAll(ctx context.Context) (products []domain.Product, err error) {
	err = p.db.Debug().WithContext(ctx).Table("products").Find(&products).Error
	return
}
