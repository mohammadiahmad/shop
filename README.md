# shop
### Simple Shop Api

This project have two api one for search between products and second for adding item to Cart.

###How to run?
1. run mysql docker image using following command <br />
   ```
   make up
   ```
2. migrate schemas using following command <br />
   ```
   go run main.go migrate
   ```
3. start api server <br />
   ```
   go run main.go server
   ```
4. tear down mysql db <br />
   ```
   make down
   ```

### Api Calls
1. search items
   ```
   curl -i -X GET 'http://127.0.0.1:9095/search?term=sh'
   ```
   
2. add product to cart (cart_item_id,cart_id is optional)
    ```
    curl -i -X POST \
       -H "Content-Type:application/json" \
       -d \
    '{
      "cart_item_id":1,
      "cart_id":1,
      "product_id":1,
      "quantity":3,
      "price":200,
      "customer_id":1
    }' \
     'http://127.0.0.1:9095/cart-item'
    ```

   
