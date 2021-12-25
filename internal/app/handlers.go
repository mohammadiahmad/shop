package app

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type SearchQuery struct {
	Term string `json:"term"`
}

type IncomingData struct {
	ProductId 	int		`json:"product_id"`
	Quantity	int 	`json:"quantity"`
}

func (a *App) search(c *fiber.Ctx) error {
	sq := &SearchQuery{}
	err := c.QueryParser(sq)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	data, err := a.db.ProductSearch(sq.Term)
	if err!=nil{
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.JSON(data)
}

func (a *App) addItemToCart(c *fiber.Ctx) error {
	cid:=c.Get("username","admin")
	data:=IncomingData{}
	err:=c.BodyParser(&data)
	if err!=nil{
		return c.SendStatus(http.StatusBadRequest)
	}

	err=a.cartStorage.AddItemToCart(c.Context(),cid,data.ProductId,data.Quantity)
	if err!=nil{
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)

}

func (a *App)deleteItemFroCart(c *fiber.Ctx) error{
	cid:=c.Get("username","admin")
	itemId,err:=c.ParamsInt("id")
	if err!=nil{
		return c.SendStatus(http.StatusBadRequest)
	}
	err=a.cartStorage.RemoveItemFromCart(c.Context(),cid,itemId)
	if err!=nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
