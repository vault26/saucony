package product

import (
	"html/template"
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/model"
	"github.com/gorilla/mux"
)

type ShowQuery struct {
	ProductId int
	Color     string
	Gender    string
	Model     string
}

type ShowProduct struct {
	ProductMap map[string]model.Product
	Query      ShowQuery
}

func Show(db database.DB) handler.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := r.URL.Query().Get("color")
		vars := mux.Vars(r)
		modelPath := vars["model_path"]
		gender := vars["gender"]

		products := db.Product(map[string]string{
			"gender":     genderShortName(gender),
			"model_path": modelPath,
		})

		productMap := make(map[string]model.Product)
		for _, product := range products {
			productMap[product.Color] = product
		}
		data := ShowProduct{
			productMap,
			ShowQuery{
				productMap[color].ID,
				color,
				gender,
				modelPath,
			},
		}

		t := template.Must(handler.BaseTemplate("product.tmpl", productsFuncMap()))
		t.ExecuteTemplate(w, "main", data)
	}
}

func genderShortName(text string) string {
	genderMap := map[string]string{"men": "M", "women": "W"}
	return genderMap[text]
}

// func (s ShowProduct) QueryProduct(attr string) string {
// 	product := s.ProductMap[s.Query.Color]
// 	r := reflect.ValueOf(product)
// 	f := reflect.Indirect(r).FieldByName(attr)
// 	if f.Kind() == reflect.Float64 {
// 		ac := accounting.Accounting{Precision: precision}
// 		return fmt.Sprint(ac.FormatMoney(number), suffix)
// 	}
// 	return f.String()
// }
