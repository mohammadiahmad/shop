package storage

import "time"

type Customer struct {
	CustomerId int    `json:"customer_id" gorm:"primary_key"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Carts      []Cart `gorm:"foreignKey:CustomerId"`
}

type ProductCategory struct {
	CategoryId   int       `json:"id" gorm:"primary_key"`
	CategoryName string    `json:"name"`
	Products     []Product `gorm:"foreignKey:Id"`
}

type ProductBrand struct {
	BrandId   int       `json:"id" gorm:"primary_key"`
	BrandName string    `json:"brand_name"`
	Products  []Product `gorm:"foreignKey:Id"`
}

type Product struct {
	Id         int        `json:"id" gorm:"primary_key"`
	Name       string     `json:"name" gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram VISIBLE"`
	BrandId    int        `json:"brand_id"`
	CategoryId int        `json:"category_id"`
	ModelYear  int        `json:"model_year"`
	Price      uint       `json:"price"`
	Quantity   uint       `json:"quantity"`
	CartItem   []CartItem `gorm:"foreignKey:CartItemId"`
}

type CartItem struct {
	CartItemId		int `json:"cart_item_id" gorm:"primaryKey"`
	CartId    		int `json:"cart_id"`
	ProductId 		int `json:"product_id"`
	Quantity  		int `json:"quantity"`
	Price     		int `json:"price"`
}

type Cart struct {
	CartId        	int        `json:"cart_id" gorm:"primary_key"`
	CustomerId    	int        `json:"customer_id"`
	PaymentStatus 	string     `json:"status"`
	CreatedAt     	time.Time  `json:"created_at"`
	CartItems     	[]CartItem `gorm:"foreignKey:CartId"`
}
