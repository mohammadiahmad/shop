package readisearch

import "time"

type Customer struct {
	CustomerId 	int    `json:"customer_id"`
	FirstName  	string `json:"first_name"`
	LastName   	string `json:"last_name"`
	Phone      	string `json:"phone"`
	Email      	string `json:"email"`
}

type Product struct {
	ProductId       int        	`json:"product_id"`
	Name       		string     	`json:"name"`
	Brand    		string      `json:"brand"`
	Category 		string      `json:"category"`
	ModelYear  		int         `json:"model_year"`
	Price      		uint       	`json:"price"`
	Quantity   		uint       	`json:"quantity"`
}

type CartItem struct {
	CartItemId		int `json:"cart_item_id"`
	CartId    		int `json:"cart_id"`
	ProductId 		int `json:"product_id"`
	Quantity  		int `json:"quantity"`
	Price     		int `json:"price"`
}

type Cart struct {
	CartId        	int        `json:"cart_id"`
	CustomerId    	int        `json:"customer_id"`
	PaymentStatus 	string     `json:"status"`
	CreatedAt     	time.Time  `json:"created_at"`
}
