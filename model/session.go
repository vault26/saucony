package model

import (
	"errors"
	"net/http"
	"strconv"

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

	data, ok := s.GetData("cart")
	// cart is empty
	if !ok {
		s.Values["cart"] = Cart{
			[]CartProduct{{*product, 1, size, product.Price}},
			product.Price,
		}
		s.Save(r, w)
		return
	}
	cart := data.(Cart)
	// add quantity to existing product
	for k, v := range cart.Products {
		if v.Product.ID == product.ID && v.Size == size {
			cart.Products[k].Quantity += 1
			cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Price
			updateOrderTotal(&cart)
			s.Save(r, w)
			return
		}
	}
	// create new product
	cart.Products = append(cart.Products, CartProduct{*product, 1, size, product.Price})
	updateOrderTotal(&cart)
	s.Values["cart"] = cart
	s.Save(r, w)
	return
}

func (s *Session) RemoveProductFromCart(
	w http.ResponseWriter,
	r *http.Request,
	productID string,
	size string) error {

	data, ok := s.GetData("cart")
	if !ok {
		return errors.New("No products in cart")
	}
	cart := data.(Cart)
	for k, v := range cart.Products {
		if productID == strconv.Itoa(v.ID) && size == v.Size {
			cart.Products = append(cart.Products[:k], cart.Products[k+1:]...)
			updateOrderTotal(&cart)
			s.Values["cart"] = cart
			s.Save(r, w)
			return nil
		}
	}
	return errors.New("No products found")
}

func (s *Session) ClearCart(w http.ResponseWriter, r *http.Request) error {
	_, ok := s.GetData("cart")
	if !ok {
		return errors.New("No products in cart")
	}
	delete(s.Values, "cart")
	s.Save(r, w)
	return nil
}

func (s *Session) AdjustOrder(
	w http.ResponseWriter,
	r *http.Request,
	productID string,
	params struct {
		Operator string
		Size     string
		Quantity int
	}) error {
	if params.Operator == "" || params.Size == "" || params.Quantity == 0 {
		return errors.New("Operator, Size, and Quantity are required")
	}
	if !(params.Operator == "add" || params.Operator == "remove") {
		return errors.New("Operator must be either 'add' or 'remove'")
	}
	data, ok := s.GetData("cart")
	if !ok {
		return errors.New("No products in cart")
	}
	cart := data.(Cart)
	for k, v := range cart.Products {
		if productID == strconv.Itoa(v.ID) && params.Size == v.Size {
			if params.Operator == "add" {
				cart.Products[k].Quantity += params.Quantity
				cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Price
			} else if params.Operator == "remove" {
				if (cart.Products[k].Quantity - params.Quantity) < 1 {
					return errors.New("Cannot remove cart product to be less than 1 item")
				}
				cart.Products[k].Quantity -= params.Quantity
				cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Price
			}
			updateOrderTotal(&cart)
			s.Values["cart"] = cart
			s.Save(r, w)
			return nil
		}
	}
	return errors.New("No products found")
}

func (s *Session) UpdateCustomerInfo(
	w http.ResponseWriter,
	r *http.Request,
	params map[string]string) {
	s.Values["customer"] = Customer{
		Firstname: params["firstname"],
		Lastname:  params["lastname"],
		Email:     params["email"],
		Phone:     params["phone"],
		Address:   params["address"],
	}
	s.Save(r, w)
}

func (s *Session) GetData(key string) (interface{}, bool) {
	data, ok := s.Values[key]
	return data, ok
}

func updateOrderTotal(cart *Cart) {
	cart.OrderTotal = 0
	for _, v := range cart.Products {
		cart.OrderTotal += v.Total
	}
}
