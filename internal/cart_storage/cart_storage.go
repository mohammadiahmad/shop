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

func NewCartStorage(cfg Config) *CartStorage {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	rc := redis.NewClient(&redis.Options{Addr: addr, DB: cfg.Db})
	cs := &CartStorage{
		redis: rc,
	}
	return cs
}

func (c *CartStorage) AddItemToCart(ctx context.Context, customer string, productId int, quantity int) error {
	_, err := c.redis.HSet(ctx, "cart:"+customer, productId, quantity).Result()
	return err
}

func (c *CartStorage) RemoveItemFromCart(ctx context.Context, customer string, productId int) error {
	_, err := c.redis.HDel(ctx, "cart:"+customer, strconv.Itoa(productId)).Result()
	return err
}
