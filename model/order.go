package model

import "time"

type Order struct {
	ID             int
	CustomerId     int
	TotalPrice     float64
	Ref            string
	CreatedAt      time.Time
	LastModifiedAt time.Time
}
