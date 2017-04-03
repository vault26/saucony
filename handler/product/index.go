package product

import (
	"net/http"
	"net/url"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type IndexHelper struct {
	ShoeSizes []string
}

type Products struct {
	model.Tmpl
	T        IndexHelper
	Genders  []string
	Sections []string
	Sizes    []string
	Types    []string
	Products []model.Product
}

func Index(w http.ResponseWriter, r *http.Request) {
	redirectWithGenders(w, r)
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	cart := ctx["cart"].(model.Cart)

	r.ParseForm()
	queryMap := make(map[string][]string)
	queryMap["genders"] = r.Form["gender[]"]
	queryMap["sections"] = r.Form["section[]"]
	queryMap["sizes"] = r.Form["size[]"]
	queryMap["types"] = r.Form["gender[]"]
	queryMap["section"] = r.Form["section[]"]
	query := r.URL.Query().Get("query")
	if query == "" {
		queryMap["query"] = r.Form["query"]
	} else {
		queryMap["query"] = []string{query}
	}

	td := &Products{
		T: IndexHelper{
			[]string{
				"5", "5.5", "6", "6.5", "7", "7.5", "8", "8.5", "9", "9.5",
				"10", "10.5", "11", "11.5", "12", "12.5", "13",
			},
		},
		Genders:  queryMap["genders"],
		Sections: queryMap["sections"],
		Sizes:    queryMap["sizes"],
		Types:    queryMap["types"],
		Products: db.Products(queryMap),
	}
	td.Title = "All Products"
	td.QueryText = query
	td.Cart = cart

	t := handler.BaseTemplate("products.tmpl", nil)
	t.ExecuteTemplate(w, "main", td)
}

func redirectWithGenders(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	if len(queryMap) == 0 {
		values := &url.Values{
			"gender[]": []string{"M", "W"},
		}
		http.Redirect(w, r, "/products?"+values.Encode(), http.StatusMovedPermanently)
	}
}
