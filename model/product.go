package model

type Product struct {
	ID    int     `gorm:"id" json:"id"`
	Name  string  `gorm:"name" json:"name"`
	Qty   int     `gorm:"qty" json:"qty"`
	Price float64 `gorm:"price" json:"price"`
}
