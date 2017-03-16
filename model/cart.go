package model

type CartProduct struct {
	Product
	Quantity int
	Size     string
}

type Cart struct {
	Products []CartProduct
}

func (c Cart) Total() (total int) {
	for _, v := range c.Products {
		total += v.Quantity
	}
	return total
}
