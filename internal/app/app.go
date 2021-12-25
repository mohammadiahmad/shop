package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mohammadiahmad/shop/internal/cart_storage"
	"github.com/mohammadiahmad/shop/internal/readisearch"
)

type App struct {
	server 		*fiber.App
	db     		*readisearch.Redisearch
	cartStorage *cart_storage.CartStorage
	config 		*Config
}

func NewApp(cfg *Config, storage *readisearch.Redisearch,cartStorage *cart_storage.CartStorage) *App {
	app := fiber.New()
	a := &App{
		server: app,
		db:     storage,
		cartStorage: cartStorage,
		config: cfg,
	}
	return a
}

func (a *App) Run() {
	a.server.Use(logger.New())
	a.server.Get("/search", a.search)
	a.server.Post("/cart-item",a.addItemToCart)
	a.server.Delete("/cart-item/:id",a.deleteItemFroCart)
	addr := fmt.Sprintf("%s:%d", a.config.Address, a.config.Port)
	err := a.server.Listen(addr)
	if err != nil {
		fmt.Println(err)
	}
}
