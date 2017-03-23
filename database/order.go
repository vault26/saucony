package database

import (
	"fmt"
	"time"

	"github.com/ekkapob/saucony/model"
)

func (db *DB) CreateOrder(customerId int, cart model.Cart) (ref string, err error) {
	order := &model.Order{
		CustomerId: customerId,
		TotalPrice: cart.OrderTotal,
	}
	err = db.Insert(order)
	if err != nil {
		return ref, err
	}
	for _, v := range cart.Products {
		orderLine := &model.OrderLine{
			OrderId:   order.ID,
			ProductId: v.ID,
			Size:      v.Size,
			Quantity:  v.Quantity,
		}
		err := db.Insert(orderLine)
		if err != nil {
			return ref, err
		}
	}
	order.Ref = orderRef(order.ID)
	_, err = db.Model(&order).Set("ref = ?ref").Where("id = ?id").Update()
	return order.Ref, err
}

func orderRef(orderId int) string {
	t := time.Now()
	return fmt.Sprintf("%02d%02d%d", t.Day(), t.Month(), orderId)
}
