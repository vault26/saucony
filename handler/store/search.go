package store

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Product struct {
	RemoteImage string
	Sizes       []float64 // need float for sorted sizes
}

type Retailer struct {
	Name       string
	Phone      string
	CityTh     string
	ProductMap map[string]Product
}

type StoresIndex struct {
	model.Tmpl
	Query       string
	Stores      []model.Store
	RetailerMap map[string]Retailer
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
	retailerMap := reduceStoreProducts(storeResults)

	tmpl := helper.InitTemplate(w, r)
	stores := &StoresIndex{
		Tmpl:        tmpl,
		Query:       queryText,
		Stores:      db.Stores(),
		RetailerMap: retailerMap,
	}

	t := handler.BaseTemplate("stores.tmpl", nil)
	t.ExecuteTemplate(w, "main", stores)
}

func reduceStoreProducts(stores []model.Store) map[string]Retailer {
	retailerMap := make(map[string]Retailer)
	for _, v := range stores {
		size, _ := strconv.ParseFloat(v.Size, 64)
		// exisiting values
		if _, ok := retailerMap[v.RetailerNo]; ok {
			retailer := retailerMap[v.RetailerNo]
			product := retailer.ProductMap[v.Model]
			if _, ok := retailer.ProductMap[v.Model]; ok {
				product.Sizes = append(product.Sizes, size)
				sort.Float64s(product.Sizes)
			} else {
				product.RemoteImage = v.RemoteImage
				product.Sizes = []float64{size}
			}
			retailer.ProductMap[v.Model] = product
			continue
		}
		// new value
		retailer := Retailer{
			Name:   v.Name,
			Phone:  v.Phone,
			CityTh: v.CityTh,
		}
		retailer.ProductMap = make(map[string]Product)
		retailer.ProductMap[v.Model] = Product{
			RemoteImage: v.RemoteImage,
			Sizes:       []float64{size},
		}
		retailerMap[v.RetailerNo] = retailer
	}
	return retailerMap
}
