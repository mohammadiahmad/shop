package app

import (
	"github.com/gofiber/fiber/v2"
)

type SearchQuery struct {
	Term string `json:"term"`
}

func (a *App) search(c *fiber.Ctx) error {
	sq := &SearchQuery{}
	err := c.QueryParser(sq)
	if err != nil {
		return err
	}
	data, err := a.db.ProductSearch(sq.Term)
	return c.JSON(data)
}

func (a *App) addItemToCart() {

}
