package delivery

import (
	"context"

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

	params := &[]model.ParamOrder{
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
					ID:    4,
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
					ID:    4,
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

	d.uc.Checkout(ctx, params)
}
