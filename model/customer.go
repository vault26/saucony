package model

import "time"

type Customer struct {
	ID             int
	Firstname      string `valid:"required"`
	Lastname       string `valid:"required"`
	Email          string `valid:"email,required"`
	Phone          string `valid:"required"`
	Address        string `valid:"required"`
	CreatedAt      time.Time
	LastModifiedAt time.Time
}
