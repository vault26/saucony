package cart

import (
	"encoding/json"
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/mux"
)

func Products(w http.ResponseWriter, r *http.Request) {
	cart := helper.GetCart(r)
	td := model.Tpl{Cart: cart}
	t := handler.BaseTemplate("", nil)
	w.Header().Set("Content-Type", "text/plain")
	t.ExecuteTemplate(w, "cart-products", td)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
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
	session.AddProductToCart(w, r, product, request.Size)
	w.WriteHeader(http.StatusOK)
}

func productSizeAvailable(product *model.Product, requestSize string) bool {
	for _, size := range product.Sizes {
		if requestSize == size {
			return true
		}
	}
	return false
}
