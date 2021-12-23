package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mohammadiahmad/shop/internal/storage"
)

type App struct {
	server *fiber.App
	db     *storage.Storage
	config *Config
}

func NewApp(cfg *Config, storage *storage.Storage) *App {
	app := fiber.New()
	a := &App{
		server: app,
		db:     storage,
		config: cfg,
	}
	return a
}

func (a *App) Run() {
	a.server.Use(logger.New())
	a.server.Get("/search", a.search)
	a.server.Post("/cart-item",a.addItemToCart)
	addr := fmt.Sprintf("%s:%d", a.config.Address, a.config.Port)
	err := a.server.Listen(addr)
	if err != nil {
		fmt.Println(err)
	}
}
