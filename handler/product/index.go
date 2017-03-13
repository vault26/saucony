package product

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/model"
)

type TemplateRender struct {
	ShoeSizes []string
}
type IndexQuery struct {
	// for template rendering
	T         TemplateRender
	Genders   []string
	Sections  []string
	Sizes     []string
	Types     []string
	QueryText string
	Products  []model.Product
}

func Index(db database.DB) handler.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexRedirect(w, r)
		query := r.URL.Query().Get("query")

		r.ParseForm()
		queryMap := make(map[string][]string)
		queryMap["genders"] = r.Form["gender[]"]
		queryMap["sections"] = r.Form["section[]"]
		queryMap["sizes"] = r.Form["size[]"]
		queryMap["types"] = r.Form["gender[]"]
		if query == "" {
			queryMap["query"] = r.Form["query"]
		} else {
			queryMap["query"] = []string{query}
		}

		t := template.Must(handler.BaseTemplate("products.tmpl", productsFuncMap()))
		templateData := productIndexTemplateData(queryMap, db.Products(queryMap))
		t.ExecuteTemplate(w, "main", templateData)
	}
}

func productIndexTemplateData(
	queryMap map[string][]string,
	products []model.Product) IndexQuery {

	shoeSizes := []string{
		"5", "5.5", "6", "6.5", "7", "7.5", "8", "8.5",
		"9", "9.5", "10", "10.5", "11", "11.5", "12", "12.5", "13",
	}
	var queryText string
	if data := queryMap["query"]; data != nil {
		queryText = data[0]
	}

	return IndexQuery{
		TemplateRender{
			shoeSizes,
		},
		queryMap["genders"],
		queryMap["sections"],
		queryMap["sizes"],
		queryMap["types"],
		queryText,
		products,
	}
}

func indexRedirect(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	if len(queryMap) == 0 {
		values := &url.Values{
			"gender[]": []string{"M", "W"},
		}
		http.Redirect(w, r, "/products?"+values.Encode(), http.StatusMovedPermanently)
	}
}
