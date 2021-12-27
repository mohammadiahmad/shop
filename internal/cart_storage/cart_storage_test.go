package cart_storage

import (
	"context"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type CartStorageTestSuite struct {
	suite.Suite
	cartStorage *CartStorage
}

func (suite *CartStorageTestSuite) SetupSuite() {
	require := suite.Require()
	server, err := miniredis.Run()
	if err != nil {
		fmt.Println("Error in run miniredis")
		panic(err)
	}
	port, _ := strconv.Atoi(server.Port())
	cfg := Config{
		Host: server.Host(),
		Port: port,
		Db:   0,
	}
	suite.cartStorage = NewCartStorage(cfg)
	require.NotNil(suite.cartStorage)
}

func (suite *CartStorageTestSuite) TestAddItemToCart() {
	require := suite.Require()
	ctx := context.Background()
	customer := "admin"
	productId := 1
	quantity := 2
	err := suite.cartStorage.AddItemToCart(ctx, customer, productId, quantity)
	require.NoError(err)

	exist, err := suite.cartStorage.redis.Exists(ctx, "cart:"+customer).Result()
	require.NoError(err)
	require.Equal(true, exist != 0)

}

func (suite *CartStorageTestSuite) TestRemoveItemFromCart() {
	require := suite.Require()
	ctx := context.Background()
	customer := "admin"
	productId := 1
	quantity := 2
	err := suite.cartStorage.AddItemToCart(ctx, customer, productId, quantity)
	require.NoError(err)

	err = suite.cartStorage.RemoveItemFromCart(ctx, customer, productId)
	require.NoError(err)

	exist, err := suite.cartStorage.redis.Exists(ctx, "cart:"+customer).Result()
	require.NoError(err)
	require.Equal(false, exist != 0)

}

func TestRedis(t *testing.T) {
	suite.Run(t, new(CartStorageTestSuite))
}
