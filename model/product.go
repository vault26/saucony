package model

import (
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/glog"
)

// Database model
type Product struct {
	ID                 int
	Model              string
	ModelPath          string
	Details            string
	Color              string
	Gender             string
	Features           []string `pg:",array"`
	Types              []string `pg:",array"`
	Subtypes           []string `pg:",array"`
	Sizes              []string `pg:",array"`
	PrimaryImage       string
	Images             []string `pg:",array"`
	PrimaryRemoteImage string
	RemoteImages       []string `pg:",array"`
	Price              float64
	Discount           float64
	Technologies       []string `pg:",array"`
	CreatedAt          time.Time
	LastModifiedAt     time.Time
}

// URL constructor for a product
// pattern: /products/{lowercase product.model}?gender={product gender}&color={product color}
// Query path is escaped
func (p Product) URL() string {
	genderMap := map[string]string{
		"M": "men",
		"W": "women",
	}
	path := fmt.Sprint("/products/", genderMap[p.Gender], "/",
		url.PathEscape(p.ModelPath))
	u, err := url.Parse(path)
	if err != nil {
		glog.Error(err)
	}
	values := url.Values{
		"color": []string{p.Color},
	}
	u.RawQuery = values.Encode()
	return u.String()
}

func (p Product) SellPrice() float64 {
	if p.Discount > 0 {
		return math.Ceil(p.Price * (1 - (p.Discount / 100)))
	}
	return p.Price
}
