package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/marketplace/internal/api"
	"github.com/khairulharu/marketplace/internal/component"
	"github.com/khairulharu/marketplace/internal/config"
	"github.com/khairulharu/marketplace/internal/repository"
	"github.com/khairulharu/marketplace/internal/service"
)

func main() {
	config := config.Get()
	dbConnetion := component.NewDatabase(config)

	userRepository := repository.NewUser(dbConnetion)
	productRepository := repository.NewProduct(dbConnetion)
	cartRepository := repository.NewCart(dbConnetion)

	cartService := service.NewCart(cartRepository)
	productService := service.NewProduct(productRepository)
	userService := service.NewUser(userRepository)

	app := fiber.New()
	// app.Use(cors.New())
	app.Use(logger.New())

	api.NewCart(cartService, app)
	api.NewUser(userService, app)
	api.NewProduct(productService, app)

	app.Listen(config.SRV.Host + ":" + config.SRV.Port)
}
