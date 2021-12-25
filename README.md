# shop
### Simple Shop Api

This project have two api one for search between products and second for adding item to Cart.

###How to run?
1. run redisearch docker image using following command <br />
   ```
   make up
   ```
2. create schemas using following command and adding test data <br />
   ```
   go run main.go migrate
   ```
3. start api server <br />
   ```
   go run main.go server
   ```
4. tear down redisearch container<br />
   ```
   make down
   ```

### Api Calls
1. search items
   ```
   curl -i -X GET 'http://127.0.0.1:9095/search?term=show'
   ```
   
2. add product to cart 
    ```
    curl -i -X POST \
       -H "Content-Type:application/json" \
       -d \
    '{
      "product_id":1,
      "quantity":3,
    }' \
     'http://127.0.0.1:9095/cart-item'
    ```
3. delete item from cart
   ```
   curl -i -X DELETE 'http://127.0.0.1:9095/cart-item/1'
   ```
   
