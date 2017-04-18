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
	size string) error {

	data, ok := s.GetData("cart")
	// cart is empty
	cartProduct := CartProduct{
		ID:                 product.ID,
		Model:              product.Model,
		PrimaryRemoteImage: product.PrimaryRemoteImage,
		Gender:             product.Gender,
		Price:              product.Price,
		Discount:           product.Discount,
		Color:              product.Color,
		Quantity:           1,
		Size:               size,
		Total:              product.SellPrice(),
	}
	if !ok {
		cart := Cart{
			Products:   []CartProduct{cartProduct},
			OrderTotal: product.SellPrice(),
		}
		s.updateOrderTotal(&cart)
		s.Values["cart"] = cart
		return s.Save(r, w)
	}
	cart := data.(Cart)
	// add quantity to existing product
	for k, v := range cart.Products {
		if v.ID == product.ID && v.Size == size {
			cart.Products[k].Quantity += 1
			cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Total
			s.updateOrderTotal(&cart)
			return s.Save(r, w)
		}
	}
	// create new product
	cart.Products = append(cart.Products, cartProduct)
	s.updateOrderTotal(&cart)
	s.Values["cart"] = cart
	return s.Save(r, w)
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
			s.updateOrderTotal(&cart)
			s.Values["cart"] = cart
			return s.Save(r, w)
		}
	}
	return errors.New("No products found")
}

func (s *Session) ClearData(w http.ResponseWriter, r *http.Request, key string) error {
	_, ok := s.GetData(key)
	if !ok {
		return errors.New("No data found")
	}
	delete(s.Values, key)
	return s.Save(r, w)
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
				cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Total
			} else if params.Operator == "remove" {
				if (cart.Products[k].Quantity - params.Quantity) < 1 {
					return errors.New("Cannot remove cart product to be less than 1 item")
				}
				cart.Products[k].Quantity -= params.Quantity
				cart.Products[k].Total = float64(cart.Products[k].Quantity) * cart.Products[k].Total
			}
			s.updateOrderTotal(&cart)
			s.Values["cart"] = cart
			return s.Save(r, w)
		}
	}
	return errors.New("No products found")
}

func (s *Session) UpdateCustomerInfo(
	w http.ResponseWriter,
	r *http.Request,
	params map[string]string) error {
	s.Values["customer"] = Customer{
		Firstname: params["firstname"],
		Lastname:  params["lastname"],
		Email:     params["email"],
		Phone:     params["phone"],
		Address:   params["address"],
	}
	return s.Save(r, w)
}

func (s *Session) GetData(key string) (interface{}, bool) {
	data, ok := s.Values[key]
	return data, ok
}

func (s *Session) ApplyPromotion(
	w http.ResponseWriter,
	r *http.Request,
	promotion Promotion) error {
	s.Values["promotion"] = promotion
	return s.Save(r, w)
}

func (s *Session) updateOrderTotal(cart *Cart) {
	cart.OrderTotal = 0
	for _, v := range cart.Products {
		cart.OrderTotal += v.Total
	}
	cart.OrderTotalBeforePromoDiscount = cart.OrderTotal
	data, ok := s.GetData("promotion")
	if !ok {
		return
	}
	promotion := data.(Promotion)
	cart.OrderTotal = cart.OrderTotal * (1 - promotion.DiscountPercent/100)
}
