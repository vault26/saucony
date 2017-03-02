package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/mux"
	"github.com/leekchan/accounting"
)

type TemplateRender struct {
	ShoeSizes []string
}
type ProductQuery struct {
	// for template rendering
	T        TemplateRender
	Genders  []string
	Sections []string
	Sizes    []string
	Types    []string
	Products []model.Product
}

func Products(db database.DB) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if Etag("products", w, r) {
			return
		}
		r.ParseForm()
		params := make(map[string][]string)
		params["sizes"] = r.Form["size[]"]
		params["genders"] = r.Form["gender[]"]

		shoeSizes := []string{
			"5", "5.5", "6", "6.5", "7", "7.5", "8", "8.5",
			"9", "9.5", "10", "10.5", "11", "11.5", "12", "12.5", "13",
		}
		t := template.Must(baseTemplate("products.tmpl", productsFuncMap()))
		data := ProductQuery{
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

func Product(db database.DB) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		product := db.Product(map[string]string{
			"color":      queries.Get("color"),
			"gender":     queries.Get("gender"),
			"model_path": mux.Vars(r)["model_path"],
		})
		_ = product
		t := template.Must(baseTemplate("product.tmpl", productsFuncMap()))
		t.ExecuteTemplate(w, "main", nil)
	}
}

func productsFuncMap() template.FuncMap {
	return template.FuncMap{
		"contains": func(list []string, text string) bool {
			for _, item := range list {
				if item == text {
					return true
				}
			}
			return false
		},
		"money": func(number float64, precision int, suffix string) string {
			ac := accounting.Accounting{Precision: precision}
			return fmt.Sprint(ac.FormatMoney(number), suffix)
		},
	}
}

func Etag(resource string, w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Etag", fmt.Sprint(`"`, resource, ":", r.URL.RawQuery, `"`))
	w.Header().Set("Cache-Control", "max-age=86400") // 1 hour (60*60*24)
	if match := r.Header.Get("If-None-Match"); match != "" {
		w.WriteHeader(http.StatusNotModified)
		return true
	}
	return false
}
