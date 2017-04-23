package store

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Product struct {
	RemoteImage string
	Gender      string
	Sizes       []float64 // need float for sorted sizes
}

type Customer struct {
	Name       string
	Phone      string
	CityTh     string
	ProductMap map[string]Product
}

type StoresIndex struct {
	model.Tmpl
	Query       string
	Stores      []model.Store
	CustomerMap map[string]Customer
}

func Stores(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	tmpl := helper.InitTemplate(w, r)
	stores := &StoresIndex{
		Tmpl:   tmpl,
		Stores: db.Stores(),
	}

	t := handler.BaseTemplate("stores.tmpl", nil)
	t.ExecuteTemplate(w, "main", stores)
}

func Search(w http.ResponseWriter, r *http.Request) {
	queryText := r.URL.Query().Get("query")
	if queryText == "" {
		http.Redirect(w, r, "/stores", http.StatusMovedPermanently)
	}
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	storeResults := db.FindStores(queryText)
	customerMap := reduceStoreProducts(storeResults)

	tmpl := helper.InitTemplate(w, r)
	stores := &StoresIndex{
		Tmpl:        tmpl,
		Query:       queryText,
		Stores:      db.Stores(),
		CustomerMap: customerMap,
	}

	t := handler.BaseTemplate("stores.tmpl", nil)
	t.ExecuteTemplate(w, "main", stores)
}

func reduceStoreProducts(stores []model.Store) map[string]Customer {
	customerMap := make(map[string]Customer)
	for _, v := range stores {
		size, _ := strconv.ParseFloat(v.Size, 64)
		// exisiting values
		if _, ok := customerMap[v.CustomerNo]; ok {
			customer := customerMap[v.CustomerNo]
			product := customer.ProductMap[v.Model]
			if _, ok := customer.ProductMap[v.Model]; ok {
				if !containSize(product.Sizes, size) {
					product.Sizes = append(product.Sizes, size)
				}
				sort.Float64s(product.Sizes)
				if product.Gender != "MW" &&
					strings.Index(product.Gender, v.Gender) == -1 {
					product.Gender = "MW"
				}
			} else {
				product.RemoteImage = v.RemoteImage
				product.Sizes = []float64{size}
			}
			customer.ProductMap[v.Model] = product
			continue
		}
		// new value
		customer := Customer{
			Name:   v.Name,
			Phone:  v.Phone,
			CityTh: v.CityTh,
		}
		customer.ProductMap = make(map[string]Product)
		customer.ProductMap[v.Model] = Product{
			RemoteImage: v.RemoteImage,
			Gender:      v.Gender,
			Sizes:       []float64{size},
		}
		customerMap[v.CustomerNo] = customer
	}
	return customerMap
}

func containSize(slice []float64, size float64) bool {
	for _, v := range slice {
		if v == size {
			return true
		}
	}
	return false
}
