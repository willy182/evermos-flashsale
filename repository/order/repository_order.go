package repository

import (
	"context"
	"fmt"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/helper"
)

// Insert method
func (r *orderRepoSQL) Insert(ctx context.Context, param *model.Order) (err error) {
	fmt.Printf("USER ID: %d | PRODUCT NAME: %s", param.UserID, param.ProductName)
	fmt.Println()
	err = r.db.Table(helper.TableOrder).Create(&param).Error
	return
}
