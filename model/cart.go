package model

type CartProduct struct {
	ID                 int
	Model              string
	PrimaryRemoteImage string
	Gender             string
	Price              float64
	Discount           float64
	Color              string
	Quantity           int
	Size               string
	Total              float64
}

type Cart struct {
	Products                      []CartProduct
	OrderTotalBeforePromoDiscount float64
	OrderTotal                    float64
}

func (c Cart) Total() (total int) {
	for _, v := range c.Products {
		total += v.Quantity
	}
	return total
}
