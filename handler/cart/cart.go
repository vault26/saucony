package cart

import (
	"encoding/json"
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/mux"
)

type Response struct {
	Error string `json:"error,omitempty"`
}

func Orders(w http.ResponseWriter, r *http.Request) {
	cart := helper.GetCart(r)
	td := model.Tpl{Cart: cart}
	t := handler.BaseTemplate("", nil)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "max-age=0")
	t.ExecuteTemplate(w, "cart-products", td)
}

func CheckoutOrders(w http.ResponseWriter, r *http.Request) {
	cart := helper.GetCart(r)
	td := model.Tpl{Cart: cart}
	t := handler.BaseTemplate("checkout.tmpl", nil)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "max-age=0")
	t.ExecuteTemplate(w, "checkout-orders", td)
}

func AddOrder(w http.ResponseWriter, r *http.Request) {
	db := helper.GetDB(r)
	session := helper.GetSession(r)
	decoder := json.NewDecoder(r.Body)
	request := struct{ Size string }{}
	decoder.Decode(&request)
	if request.Size == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if session == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	productID := vars["product_id"]
	product, err := db.Product(map[string]string{"id": productID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !productSizeAvailable(product, request.Size) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = session.AddProductToCart(w, r, product, request.Size)
	if err != nil {
		formatError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	session := helper.GetSession(r)
	size := r.URL.Query().Get("size")
	vars := mux.Vars(r)
	productID := vars["product_id"]
	err := session.RemoveProductFromCart(w, r, productID, size)
	if err != nil {
		formatError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AdjustOrder(w http.ResponseWriter, r *http.Request) {
	session := helper.GetSession(r)
	decoder := json.NewDecoder(r.Body)
	request := struct {
		Operator string
		Size     string
		Quantity int
	}{}
	decoder.Decode(&request)
	vars := mux.Vars(r)
	productID := vars["product_id"]

	err := session.AdjustOrder(w, r, productID, request)
	if err != nil {
		formatError(w, err)
		return
	}
}

func productSizeAvailable(product *model.Product, requestSize string) bool {
	for _, size := range product.Sizes {
		if requestSize == size {
			return true
		}
	}
	return false
}

func formatError(w http.ResponseWriter, err error) {
	js, _ := json.Marshal(Response{Error: err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(js)
}
