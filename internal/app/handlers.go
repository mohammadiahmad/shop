package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mohammadiahmad/shop/internal/storage"
	"net/http"
)

type SearchQuery struct {
	Term string `json:"term"`
}

type IncomingData struct {
	*storage.CartItem
	CustomerID int `json:"customer_id"`
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
	ci:=&storage.CartItem{}
	cart:=&storage.Cart{}
	data:=IncomingData{
		ci,
		0,
	}
	err:=c.BodyParser(&data)
	if err!=nil{
		fmt.Println(err)
		return c.SendStatus(http.StatusBadRequest)
	}

	if data.CartItem.CartId!=0{
		d,err:=a.db.AddCartItem(data.CartItem)
		if err!=nil{
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(d)
	} else{
		cart.CustomerId=data.CustomerID
		cart.PaymentStatus="not_payed"
		cartId,err:=a.db.CreateCart(cart)
		if err!=nil{
			return c.SendStatus(http.StatusInternalServerError)
		}
		data.CartItem.CartId=cartId
		d,err:=a.db.AddCartItem(data.CartItem)
		if err!=nil{
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(d)

	}

}
