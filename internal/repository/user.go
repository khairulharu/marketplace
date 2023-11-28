package repository

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	err = u.db.Debug().WithContext(ctx).Table("users").Where("id=?", id).First(&user).Error
	return
}

func (u userRepository) Insert(ctx context.Context, user *domain.User) error {
	err := u.db.Debug().WithContext(ctx).Table("users").Save(user).Error
	return err
}
