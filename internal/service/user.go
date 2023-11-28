package service

import (
	"context"

	"github.com/khairulharu/marketplace/domain"
	"github.com/khairulharu/marketplace/dto"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) UserRegister(ctx context.Context, req dto.UserReq) dto.Response {
	var user = domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := u.userRepository.Insert(ctx, &user); err != nil {
		return dto.Response{
			Code:    "401",
			Massage: "INVALID",
			Error:   err.Error(),
			Data:    nil,
		}
	}

	return dto.Response{
		Code:    "200",
		Massage: "APPROVE",
	}
}
