package repository

import (
	"context"
	"fmt"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/helper"
	"gorm.io/gorm/clause"
)

// Insert method
func (r *orderRepoSQL) Insert(ctx context.Context, param *model.Order) (err error) {
	opName := "RepoOrder-Insert"
	defer func() {
		if rec := recover(); rec != nil {
			msg := fmt.Sprintf("scope: %s | panic: %v", opName, rec)
			err = fmt.Errorf(msg)
		}
	}()

	err = r.db.Clauses(clause.Locking{Strength: "INSERT"}).Table(helper.TableOrder).Create(&param).Error
	if err != nil {
		return
	}

	return
}
