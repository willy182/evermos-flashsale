package model

import "time"

// Order data of struct
type Order struct {
	ID          int       `gorm:"id" json:"id"`
	OrderID     string    `gorm:"order_id" json:"order_id"`
	UserID      int       `gorm:"user_id" json:"user_id"`
	ProductName string    `gorm:"product_name" json:"product_name"`
	Price       float64   `gorm:"price" json:"price"`
	Qty         int       `gorm:"qty" json:"qty"`
	TotalPrice  float64   `gorm:"total_price" json:"total_price"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
}

// ParamOrder data of struct
type ParamOrder struct {
	UserID int    `json:"user_id"`
	Cart   []Cart `json:"cart"`
}

// Cart data of struct
type Cart struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}
