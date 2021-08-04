package repository

import (
	"context"

	"github.com/willy182/evermos-flashsale/model"
	"gorm.io/gorm"
)

type orderRepoSQL struct {
	db *gorm.DB
}

// NewOrderRepoSQL mongo repo constructor
func NewOrderRepoSQL(db *gorm.DB) OrderRepository {
	return &orderRepoSQL{
		db,
	}
}

// OrderRepository abstract interface
type OrderRepository interface {
	Insert(ctx context.Context, param *model.Order) (err error)
}
