package model

import "time"

// Order data of struct
type Order struct {
	OrderTrx    string    `gorm:"order_trx" json:"order_trx"`
	UserID      int       `gorm:"user_id" json:"user_id"`
	ProductName string    `gorm:"product_name" json:"product_name"`
	Price       int       `gorm:"price" json:"price"`
	Qty         int       `gorm:"qty" json:"qty"`
	TotalPrice  int       `gorm:"total_price" json:"total_price"`
	CreatedAt   time.Time `gorm:"created_at" json:"created_at"`
}

// ParamOrder data of struct
type ParamOrder struct {
	UserID int    `json:"user_id"`
	Cart   []Cart `json:"cart"`
}

// Cart data of struct
type Cart struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}
