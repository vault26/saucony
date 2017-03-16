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
func GetSession(r *http.Request) *model.Session {
	session, _ := r.Context().Value("session").(*model.Session)
	return session
}
