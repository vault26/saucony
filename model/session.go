package model

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Session struct {
	*sessions.Session
}

func (s *Session) AddProductToCart(
	w http.ResponseWriter,
	r *http.Request,
	product *Product,
	size string) {

	cart, ok := s.Values["cart"].(Cart)
	// cart is empty
	if !ok {
		s.Values["cart"] = Cart{
			[]CartProduct{{*product, 1, size}},
		}
		s.Save(r, w)
		return
	}
	// add quantity to existing product
	for k, v := range cart.Products {
		if v.Product.ID == product.ID && v.Size == size {
			cart.Products[k].Quantity += 1
			s.Save(r, w)
			return
		}
	}
	// create new product
	cart.Products = append(cart.Products, CartProduct{*product, 1, size})
	s.Values["cart"] = cart
	s.Save(r, w)
	return
}
