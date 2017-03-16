package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/handler/cart"
	mw "github.com/ekkapob/saucony/handler/middleware"
	"github.com/ekkapob/saucony/handler/page"
	"github.com/ekkapob/saucony/handler/product"
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
	r.HandleFunc("/", mw.PublicPage(publicParams, page.Home))

	r.HandleFunc("/products", mw.PublicPage(publicParams, product.Index))
	r.HandleFunc("/products/{gender:men|women}", product.Gender)
	r.HandleFunc("/products/{gender:men|women}/{model_path}",
		mw.PublicPage(publicParams, product.Show))

	r.HandleFunc("/cart/products/{product_id}",
		mw.PublicPage(publicParams, cart.AddProduct)).Methods("POST")
	r.HandleFunc("/cart/products", mw.PublicPage(publicParams, cart.Products))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	portname := fmt.Sprintf(":%v", *port)
	fmt.Println("Server is running on", portname)
	log.Fatal(http.ListenAndServe(portname, loggedRouter))
}
