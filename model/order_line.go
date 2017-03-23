package model

import "time"

type OrderLine struct {
	ID             int
	OrderId        int
	ProductId      int
	Size           string
	Quantity       int
	CreatedAt      time.Time
	LastModifiedAt time.Time
}
