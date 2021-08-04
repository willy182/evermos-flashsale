package model

import "time"

type Order struct {
	ID          int       `gorm:"id" json:"id"`
	OrderID     string    `gorm:"order_id" json:"order_id"`
	UserID      string    `gorm:"user_id" json:"user_id"`
	ProductName string    `gorm:"product_name" json:"product_name"`
	Price       float64   `gorm:"price" json:"price"`
	Qty         int       `gorm:"qty" json:"qty"`
	TotalPrice  float64   `gorm:"total_price" json:"total_price"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
}
