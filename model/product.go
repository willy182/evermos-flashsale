package model

// Product data of struct
type Product struct {
	ID    int    `gorm:"id" json:"id"`
	Name  string `gorm:"name" json:"name"`
	Qty   int    `gorm:"qty" json:"qty"`
	Price int    `gorm:"price" json:"price"`
}
