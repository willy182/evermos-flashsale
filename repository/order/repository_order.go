package repository

import (
	"context"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/helper"
	"gorm.io/gorm/clause"
)

// Insert method
func (r *orderRepoSQL) Insert(ctx context.Context, param *model.Order) (err error) {
	err = r.db.Clauses(clause.Locking{Strength: "INSERT"}).Table(helper.TableOrder).Create(&param).Error
	return
}
