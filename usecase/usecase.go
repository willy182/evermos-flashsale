package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/pkg/shared"
	"github.com/willy182/evermos-flashsale/repository"
)

// IUsecase
type IUsecase interface {
	Checkout(ctx context.Context, params model.ParamOrder) error
}

// orderUsecase
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
func (uc *orderUsecase) Checkout(ctx context.Context, param model.ParamOrder) (err error) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("panic: %v", r)
			log.Println(msg)
			err = fmt.Errorf(msg)
		}
	}()

	var paramOrder model.Order
	now := time.Now()
	orderTrx := shared.GenerateOrderID()
	userID := param.UserID
	for _, cart := range param.Cart {
		if err := uc.repo.WithTransaction(ctx, func(ctx context.Context, manager *repository.RepoSQL) error {
			qty := cart.Qty
			price := cart.Price
			totalPrice := qty * price

			paramOrder.OrderTrx = orderTrx
			paramOrder.UserID = userID
			paramOrder.ProductName = cart.Name
			paramOrder.Price = price
			paramOrder.Qty = qty
			paramOrder.TotalPrice = totalPrice
			paramOrder.CreatedAt = now

			cartID := cart.ID
			e := manager.ProductRepo.Update(ctx, cartID, qty)
			if e != nil {
				log.Println(e.Error())
				return e
			}

			e = manager.OrderRepo.Insert(ctx, &paramOrder)
			if e != nil {
				log.Println(e.Error())
				return e
			}

			return nil
		}); err != nil {
			panic(err)
		}
	}

	return
}
