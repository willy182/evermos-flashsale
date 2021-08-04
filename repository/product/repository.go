package repository

import (
	"context"

	"github.com/willy182/evermos-flashsale/model"
	"gorm.io/gorm"
)

type productRepoSQL struct {
	db *gorm.DB
}

// NewProductRepoSQL mongo repo constructor
func NewProductRepoSQL(db *gorm.DB) ProductRepository {
	return &productRepoSQL{
		db,
	}
}

// ProductRepository abstract interface
type ProductRepository interface {
	FindByID(ctx context.Context, id int) (result model.Product, err error)
	Update(ctx context.Context, id, qty int) (err error)
}
