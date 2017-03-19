package model

type CartProduct struct {
	Product
	Quantity int
	Size     string
	Total    float64
}

type Cart struct {
	Products   []CartProduct
	OrderTotal float64
}

func (c Cart) Total() (total int) {
	for _, v := range c.Products {
		total += v.Quantity
	}
	return total
}
