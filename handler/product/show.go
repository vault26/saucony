package product

import (
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/mux"
)

type ShowQuery struct {
	ProductId int
	Color     string
	Gender    string
	Model     string
}

type Product struct {
	model.Tmpl
	ProductMap    map[string]model.Product
	Query         ShowQuery
	AlreadyInCart bool
}

func Show(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	cart := ctx["cart"].(model.Cart)
	promotion := ctx["promotion"].(model.Promotion)

	color := r.URL.Query().Get("color")
	vars := mux.Vars(r)
	modelPath := vars["model_path"]
	gender := vars["gender"]

	products := db.ModelProducts(map[string]string{
		"gender":     genderShortName(gender),
		"model_path": modelPath,
	})
	if len(products) == 0 {
		http.Redirect(w, r, "/products", http.StatusMovedPermanently)
	}

	productMap := make(map[string]model.Product)
	firstColor := ""
	queryColorMatch := false
	for _, product := range products {
		productMap[product.Color] = product
		if firstColor == "" {
			firstColor = product.Color
		}
		if !queryColorMatch {
			queryColorMatch = (color == product.Color)
		}
	}

	if color == "" || !queryColorMatch {
		color = firstColor
	}
	productId := productMap[color].ID

	td := &Product{
		ProductMap: productMap,
		Query: ShowQuery{
			productId,
			color,
			gender,
			modelPath,
		},
	}
	td.Title = productMap[color].Model
	td.Cart = cart
	td.Promotion = promotion
	for _, v := range cart.Products {
		if productId == v.ID {
			td.AlreadyInCart = true
		}
	}

	t := handler.BaseTemplate("product.tmpl", nil)
	t.ExecuteTemplate(w, "main", td)
}

func genderShortName(text string) string {
	genderMap := map[string]string{"men": "M", "women": "W"}
	return genderMap[text]
}
