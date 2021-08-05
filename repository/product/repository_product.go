package repository

import (
	"context"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/helper"
	"gorm.io/gorm/clause"
)

// FindByID method
func (r *productRepoSQL) FindByID(ctx context.Context, id int) (result model.Product, err error) {
	err = r.db.Table(helper.TableProduct).Where(`"id" = ?`, id).First(&result).Error
	return
}

// Update method
func (r *productRepoSQL) Update(ctx context.Context, id, qty int) (err error) {
	product := &model.Product{}
	err = r.db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).Where("qty >= ?", qty).First(&product).Error
	if err != nil {
		return
	}

	product.Qty -= qty
	err = r.db.Table(helper.TableProduct).Where("id = ?", id).Update("qty", product.Qty).Error
	if err != nil {
		return
	}

	return
}
