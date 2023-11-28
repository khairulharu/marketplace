package domain

import (
	"context"

	"github.com/khairulharu/marketplace/dto"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

type UserRepository interface {
	Insert(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id int64) (User, error)
}

type UserService interface {
	UserRegister(ctx context.Context, req dto.UserReq) dto.Response
}
