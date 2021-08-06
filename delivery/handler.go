package delivery

import (
	"context"
	"log"
	"sync"

	"github.com/willy182/evermos-flashsale/model"
	"github.com/willy182/evermos-flashsale/usecase"
)

// IDelivery
type IDelivery struct {
	uc usecase.IUsecase
}

// NewDelivery constructor
func NewDelivery(uc usecase.IUsecase) *IDelivery {
	return &IDelivery{
		uc: uc,
	}
}

func (d *IDelivery) Checkout() {
	ctx := context.Background()

	params := []model.ParamOrder{
		{
			UserID: 1,
			Cart: []model.Cart{
				{
					ID:    1,
					Name:  "laptop",
					Price: 15000000,
					Qty:   2,
				},
				{
					ID:    4,
					Name:  "selai coklat",
					Price: 300000,
					Qty:   10,
				},
			},
		},
		{
			UserID: 5,
			Cart: []model.Cart{
				{
					ID:    1,
					Name:  "laptop",
					Price: 15000000,
					Qty:   2,
				},
				{
					ID:    2,
					Name:  "gamis",
					Price: 100000,
					Qty:   2,
				},
				{
					ID:    3,
					Name:  "lipstik",
					Price: 35000,
					Qty:   1,
				},
			},
		},
		{
			UserID: 3,
			Cart: []model.Cart{
				{
					ID:    1,
					Name:  "laptop",
					Price: 15000000,
					Qty:   2,
				},
			},
		},
		{
			UserID: 4,
			Cart: []model.Cart{
				{
					ID:    3,
					Name:  "lipstik",
					Price: 35000,
					Qty:   3,
				},
				{
					ID:    2,
					Name:  "gamis",
					Price: 100000,
					Qty:   3,
				},
			},
		},
		{
			UserID: 2,
			Cart: []model.Cart{
				{
					ID:    4,
					Name:  "selai coklat",
					Price: 300000,
					Qty:   5,
				},
			},
		},
	}

	var wg sync.WaitGroup
	errChan := make(chan error)

	for _, val := range params {
		var (
			payload model.ParamOrder
			carts   []model.Cart
		)

		payload.UserID = val.UserID
		for _, c := range val.Cart {
			var cart model.Cart
			cart.ID = c.ID
			cart.Name = c.Name
			cart.Price = c.Price
			cart.Qty = c.Qty
			carts = append(carts, cart)
		}

		payload.Cart = carts

		wg.Add(1)
		go func() {
			defer wg.Done()
			errChan <- d.uc.Checkout(ctx, payload)
		}()
	}

	go func() {
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		if err != nil {
			log.Println("handler error", err)
		}
	}
}
