package middleware

import (
	"context"
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/sessions"
)

func PublicPage(
	params map[string]interface{},
	next handler.HandleFunc) handler.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := params["database"].(database.DB)
		cookieStore := params["cookie-store"].(*sessions.CookieStore)

		ctx := context.WithValue(r.Context(), "db", db)
		ctx = withSessionCtx(cookieStore, r.WithContext(ctx))
		ctx = withSessionDataCtx(r.WithContext(ctx), []string{"cart", "customer"})
		next(w, r.WithContext(ctx))
	}
}

func withSessionCtx(store *sessions.CookieStore, r *http.Request) context.Context {
	session, err := store.Get(r, "sauconythailand")
	if err != nil {
		return r.Context()
	}
	return context.WithValue(r.Context(), "session", &model.Session{session})
}

func withSessionDataCtx(r *http.Request, names []string) context.Context {
	session, ok := r.Context().Value("session").(*model.Session)
	if !ok {
		return r.Context()
	}
	ctx := r.Context()
	for _, v := range names {
		data, ok := session.Values[v]
		if ok {
			ctx = context.WithValue(ctx, v, data)
		}
	}
	return ctx
}

func withCartCtx(r *http.Request) context.Context {
	session, ok := r.Context().Value("session").(*model.Session)
	if !ok {
		return r.Context()
	}
	cart, ok := session.Values["cart"].(model.Cart)
	if !ok {
		return r.Context()
	}
	return context.WithValue(r.Context(), "cart", cart)
}

func withCustomerCtx(r *http.Request) context.Context {
	session, ok := r.Context().Value("session").(*model.Session)
	if !ok {
		return r.Context()
	}
	customer, ok := session.Values["customer"].(model.Customer)
	if !ok {
		return r.Context()
	}
	return context.WithValue(r.Context(), "customer", customer)
}
