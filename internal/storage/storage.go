package storage

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewDB(cfg Config, logger *zap.Logger) (*Storage, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.Error("Error in initiate db connection", zap.Error(err))
		return nil, err
	}
	s := &Storage{
		db:     db,
		logger: logger,
	}
	return s, nil
}

func (s *Storage) Migrate() error {
	return s.db.AutoMigrate(&Customer{}, &ProductCategory{}, &ProductBrand{}, &Product{}, &Cart{}, &CartItem{})
}

//InitDB	This function insert test data to database
func (s *Storage) InitDB() {
	costumer := &Customer{
		FirstName: "Ahmad",
		LastName:  "M",
		Email:     "a@b.com",
		Phone:     "+989190000000",
	}
	s.db.Create(costumer)

	productCategory := &ProductCategory{
		CategoryName: "sport",
	}
	s.db.Create(productCategory)

	productBrand := &ProductBrand{
		BrandName: "adidas",
	}
	s.db.Create(productBrand)

	product := &Product{
		Name:       "shows",
		CategoryId: productCategory.CategoryId,
		BrandId:    productBrand.BrandId,
		ModelYear:  2021,
		Price:      150,
		Quantity:   10,
	}
	s.db.Create(product)

}

func (s *Storage) ProductSearch(term string) ([]Product, error) {
	var result []Product

	err := s.db.Table("product").Where("WHERE MATCH (name) AGAINST (? IN BOOLEAN MODE)", term).Find(&result).Error
	if err != nil {
		fmt.Printf("Error in fetching search result: %s\n", err)
		return nil, err
	}
	return result, err
}

func (s *Storage) CreateCart() {

}

func (s *Storage) AddCartItem() {

}
