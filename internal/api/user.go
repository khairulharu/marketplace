package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/marketplace/domain"
	"github.com/khairulharu/marketplace/dto"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(userService domain.UserService, app *fiber.App) {
	h := userApi{
		userService: userService,
	}
	app.Post("/register/user", h.RegisterUser)
}

func (u userApi) RegisterUser(ctx *fiber.Ctx) error {
	var reqUser dto.UserReq
	if err := ctx.BodyParser(&reqUser); err != nil {
		return ctx.SendStatus(400)
	}
	res := u.userService.UserRegister(ctx.Context(), reqUser)

	return ctx.Status(200).JSON(res)
}
