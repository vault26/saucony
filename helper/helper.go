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
	if flashes := session.Flashes(); len(flashes) > 0 {
		flash = flashes[0].(model.Flash)
		session.Save(r, w)
	}
	return flash
}
