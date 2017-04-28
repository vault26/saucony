package helper

import (
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/model"
)

func GetDB(r *http.Request) database.DB {
	db, _ := r.Context().Value("db").(database.DB)
	return db
}
func GetCart(r *http.Request) model.Cart {
	cart, _ := r.Context().Value("cart").(model.Cart)
	return cart
}
func GetCustomer(r *http.Request) model.Customer {
	customer, _ := r.Context().Value("customer").(model.Customer)
	return customer
}
func GetSession(r *http.Request) *model.Session {
	session, _ := r.Context().Value("session").(*model.Session)
	return session
}

func GetFlash(w http.ResponseWriter, r *http.Request) (flash model.Flash) {
	session := GetSession(r)
	if session == nil {
		return flash
	}
	if flashes := session.Flashes(); len(flashes) > 0 {
		flash = flashes[0].(model.Flash)
		session.Save(r, w)
	}
	return flash
}

func GetPromotion(r *http.Request) model.Promotion {
	promotion, _ := r.Context().Value("promotion").(model.Promotion)
	return promotion
}

func GetContext(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	return map[string]interface{}{
		"db":        GetDB(r),
		"session":   GetSession(r),
		"cart":      GetCart(r),
		"customer":  GetCustomer(r),
		"flash":     GetFlash(w, r),
		"promotion": GetPromotion(r),
	}
}

func InitTemplate(w http.ResponseWriter, r *http.Request) model.Tmpl {
	ctx := GetContext(w, r)
	return model.Tmpl{
		Flash:     ctx["flash"].(model.Flash),
		Cart:      ctx["cart"].(model.Cart),
		Customer:  ctx["customer"].(model.Customer),
		Promotion: ctx["promotion"].(model.Promotion),
	}
}
