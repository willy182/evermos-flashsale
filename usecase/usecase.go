package usecase

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/shared"
	"github.com/willy182/evermos-flashsale/repository"
)

type IUsecase interface {
	Checkout(ctx context.Context, params *[]model.ParamOrder) (err error)
}

type orderUsecase struct {
	repo *repository.RepoSQL
}

// NewUseCase func
func NewUseCase(repo *repository.RepoSQL) IUsecase {
	return &orderUsecase{
		repo: repo,
	}
}

// Insert method
func (uc *orderUsecase) Checkout(ctx context.Context, params *[]model.ParamOrder) (err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("panic: %v", r)
			err = fmt.Errorf(msg)
		}
	}()

	var (
		wg         sync.WaitGroup
		paramOrder model.Order
	)

	errChan := make(chan error)

	for _, val := range *params {
		if err = uc.repo.WithTransaction(ctx, func(ctx context.Context, manager *repository.RepoSQL) error {
			now := time.Now()
			paramOrder.OrderID = shared.GenerateOrderID(10)
			paramOrder.UserID = val.UserID
			paramOrder.CreatedAt = now
			for _, cart := range val.Cart {
				qty := cart.Qty
				price := cart.Price
				totalPrice := float64(qty) * price
				paramOrder.ProductName = cart.Name
				paramOrder.Price = price
				paramOrder.Qty = qty
				paramOrder.TotalPrice = totalPrice

				wg.Add(1)
				go func(id, qty int) {
					defer wg.Done()
					e := manager.ProductRepo.Update(ctx, id, qty)
					if e != nil {
						errChan <- e
					}

					e = manager.OrderRepo.Insert(ctx, &paramOrder)
					if e != nil {
						errChan <- e
					}
				}(cart.ID, cart.Qty)
			}

			return nil
		}); err != nil {
			panic(err)
		}
	}

	go func() {
		defer close(errChan)
		wg.Wait()
	}()

	// get offer data from buffered channel offerChan
	for e := range errChan {
		if e != nil {
			err = e
			return
		}
	}

	return
}
