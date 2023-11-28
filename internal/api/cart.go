package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/marketplace/domain"
	"github.com/khairulharu/marketplace/dto"
)

type cartApi struct {
	cartService domain.CartService
}

func NewCart(cartService domain.CartService, app *fiber.App) {
	h := cartApi{
		cartService: cartService,
	}

	app.Post("cart/add", h.AddProductToCart)
	app.Get("cart/get", h.GetUserCarts)
}

func (c cartApi) AddProductToCart(ctx *fiber.Ctx) error {
	var req dto.CartReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	res := c.cartService.AddProduct(ctx.Context(), req)

	return ctx.Status(200).JSON(res)
}

func (c cartApi) GetUserCarts(ctx *fiber.Ctx) error {
	var req dto.UserCartReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}
	res := c.cartService.GetCart(ctx.Context(), req)
	return ctx.Status(200).JSON(res)
}
