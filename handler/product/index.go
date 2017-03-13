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
	T        TemplateRender
	Genders  []string
	Sections []string
	Sizes    []string
	Types    []string
	Products []model.Product
}

func Index(db database.DB) handler.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		indexRedirect(w, r)

		r.ParseForm()
		params := make(map[string][]string)
		params["sizes"] = r.Form["size[]"]
		params["genders"] = r.Form["gender[]"]

		shoeSizes := []string{
			"5", "5.5", "6", "6.5", "7", "7.5", "8", "8.5",
			"9", "9.5", "10", "10.5", "11", "11.5", "12", "12.5", "13",
		}
		t := template.Must(handler.BaseTemplate("products.tmpl", productsFuncMap()))
		data := IndexQuery{
			TemplateRender{
				shoeSizes,
			},
			params["genders"],
			r.Form["section[]"],
			params["sizes"],
			r.Form["type[]"],
			db.Products(params),
		}
		t.ExecuteTemplate(w, "main", data)
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
