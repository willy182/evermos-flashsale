package repository

import (
	"context"
	"fmt"

	repoOrder "github.com/willy182/evermos-flashsale/repository/order"
	repoProduct "github.com/willy182/evermos-flashsale/repository/product"
	"gorm.io/gorm"
)

// RepoSQL abstraction
type RepoSQL struct {
	db *gorm.DB

	// register all repository from modules
	OrderRepo   repoOrder.OrderRepository
	ProductRepo repoProduct.ProductRepository
}

// NewRepositorySQL constructor
func NewRepositorySQL(db *gorm.DB) *RepoSQL {
	return &RepoSQL{
		db:          db,
		OrderRepo:   repoOrder.NewOrderRepoSQL(db),
		ProductRepo: repoProduct.NewProductRepoSQL(db),
	}
}

// WithTransaction run transaction for each repository with context, include handle canceled or timeout context
func (r *RepoSQL) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, repo *RepoSQL) error) (err error) {
	tx := r.db.Begin()
	err = tx.Error
	if err != nil {
		return err
	}

	// reinit new repository in different memory address with tx value
	manager := NewRepositorySQL(tx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
		manager.free()
	}()

	errChan := make(chan error)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic: %v", r)
			}
			close(errChan)
		}()

		if err := txFunc(ctx, manager); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("Canceled or timeout: %v", ctx.Err())
	case e := <-errChan:
		return e
	}
}

// free
func (r *RepoSQL) free() {
	// make nil all repository
	r.OrderRepo = nil
	r.ProductRepo = nil
}
