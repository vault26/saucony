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
		ctx = withCartCtx(r.WithContext(ctx))
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
