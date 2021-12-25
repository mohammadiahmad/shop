package cart_storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type CartStorage struct {
	redis *redis.Client
}



func NewCartStorage(cfg Config)  *CartStorage{
	addr:= fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	rc:=redis.NewClient(&redis.Options{Addr: addr,DB: cfg.Db})
	cs:=&CartStorage{
		redis: rc,
	}
	return cs
}

func (c *CartStorage) AddItemToCart(ctx context.Context,customer string,product_id int,quantity int) error {
	_,err:=c.redis.HSet(ctx,"cart:"+customer,product_id,quantity).Result()
	fmt.Println(err)
	return err
}

func (c *CartStorage) RemoveItemFromCart(ctx context.Context,customer string,product_id int) error {
	_,err:=c.redis.HDel(ctx,"cart:"+customer,strconv.Itoa(product_id)).Result()
	return err
}
