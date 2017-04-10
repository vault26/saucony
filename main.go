package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/handler/cart"
	"github.com/ekkapob/saucony/handler/checkout"
	mw "github.com/ekkapob/saucony/handler/middleware"
	"github.com/ekkapob/saucony/handler/order"
	"github.com/ekkapob/saucony/handler/page"
	"github.com/ekkapob/saucony/handler/product"
	"github.com/ekkapob/saucony/handler/store"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/namsral/flag"
	"gopkg.in/pg.v5"
)

func main() {
	db := database.DB{
		pg.Connect(&pg.Options{
			User:     "saucony_admin",
			Password: "sauconyrocks",
			Database: "saucony",
		}),
	}
	defer db.Close()
	publicParams := map[string]interface{}{
		"database":     db,
		"cookie-store": sessions.NewCookieStore([]byte("an-awesome-website")),
	}

	port := flag.Int("port", 8080, "port")
	flag.Parse()

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	// Page
	r.HandleFunc("/", mw.PublicPage(publicParams, page.Home))
	r.HandleFunc("/history", mw.PublicPage(publicParams, page.History))
	r.HandleFunc("/technology", mw.PublicPage(publicParams, page.Technology))
	r.HandleFunc("/stores", mw.PublicPage(publicParams, store.Stores))
	r.HandleFunc("/sale", mw.PublicPage(publicParams, page.Sale))

	// Stores
	r.HandleFunc("/stores/search", mw.PublicPage(publicParams, store.Search))

	// Products
	r.HandleFunc("/products", mw.PublicPage(publicParams, product.Index))
	r.HandleFunc("/products/{gender:men|women}", product.Gender)
	r.HandleFunc("/products/{gender:men|women}/{model_path}",
		mw.PublicPage(publicParams, product.Show))

	// Cart
	r.HandleFunc("/cart/orders/{product_id}",
		mw.PublicPage(publicParams, cart.AddOrder)).Methods("POST")
	r.HandleFunc("/cart/orders/{product_id}",
		mw.PublicPage(publicParams, cart.DeleteOrder)).Methods("DELETE")
	r.HandleFunc("/cart/orders/{product_id}",
		mw.PublicPage(publicParams, cart.AdjustOrder)).Methods("PUT")
	r.HandleFunc("/cart/orders", mw.PublicPage(publicParams, cart.Orders))
	r.HandleFunc("/cart/checkout_orders", mw.PublicPage(publicParams, cart.CheckoutOrders))

	// Checkout
	r.HandleFunc("/checkout", mw.PublicPage(publicParams, checkout.Index))

	// Order
	r.HandleFunc("/order", mw.PublicPage(publicParams, order.Create)).
		Methods("POST")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	portname := fmt.Sprintf(":%v", *port)
	fmt.Println("Server is running on", portname)
	log.Fatal(http.ListenAndServe(portname, loggedRouter))
}
