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
	Model            string
	Color            string
	Size             float64
	Stores           []model.Store
	CustomerMap      map[string]Customer
	ProductOptionMap map[string][]string
	ShoeSizes        []float64
}

func Stores(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	tmpl := helper.InitTemplate(w, r)
	stores := &StoresIndex{
		Tmpl:             tmpl,
		Stores:           db.Stores(),
		ProductOptionMap: db.ProductOptions(),
		ShoeSizes:        shoeSizes(),
	}

	t := handler.BaseTemplate("stores.tmpl", nil)
	t.ExecuteTemplate(w, "main", stores)
}

func Search(w http.ResponseWriter, r *http.Request) {
	model := r.URL.Query().Get("model")
	color := r.URL.Query().Get("color")
	size := r.URL.Query().Get("size")

	if model == "" {
		http.Redirect(w, r, "/stores", http.StatusMovedPermanently)
	}
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	storeResults := db.FindStores(model, color, size)
	customerMap := reduceStoreProducts(storeResults)

	tmpl := helper.InitTemplate(w, r)
	floatSize, _ := strconv.ParseFloat(size, 64)
	stores := &StoresIndex{
		Tmpl:             tmpl,
		Model:            model,
		Color:            color,
		Size:             floatSize,
		Stores:           db.Stores(),
		ProductOptionMap: db.ProductOptions(),
		ShoeSizes:        shoeSizes(),
		CustomerMap:      customerMap,
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

func shoeSizes() []float64 {
	return []float64{
		5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9,
		9.5, 10, 10.5, 11, 11.5, 12, 12.5, 13}
}
