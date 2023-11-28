package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/marketplace/domain"
)

type productApi struct {
	productService domain.ProductService
}

func NewProduct(productService domain.ProductService, app *fiber.App) {
	h := productApi{
		productService: productService,
	}

	app.Get("products/get", h.GetProducts)
}

func (p productApi) GetProducts(ctx *fiber.Ctx) error {
	res := p.productService.ShowProduct(ctx.Context())
	return ctx.Status(200).JSON(res)
}
