package database

import (
	"github.com/ekkapob/saucony/model"
	"github.com/golang/glog"
	pg "gopkg.in/pg.v5"
)

func (db *DB) Products(params map[string][]string) (products []model.Product) {
	query := db.Model(&products)
	if genders, ok := paramsValue(params, "genders"); ok {
		query.Where("gender IN (?)", pg.In(genders))
	}
	if sizes, ok := paramsValue(params, "sizes"); ok {
		query.Where("sizes && ?", pg.Array(sizes))
	}
	err := query.Order("model").Select()
	if err != nil {
		glog.Error(err)
	}
	return products
}

func (db *DB) Product(params map[string]string) (product model.Product) {
	query := db.Model(&product)
	if params["color"] != "" {
		query.Where("color = ?", params["color"])
	}
	if params["gender"] != "" {
		query.Where("gender = ?", params["gender"])
	}
	err := query.Where("model_path = ?", params["model_path"]).First()
	if err != nil {
		glog.Error(err)
	}
	return product
}

func paramsValue(params map[string][]string, key string) ([]string, bool) {
	if value := params[key]; len(value) > 0 {
		return value, true
	}
	return nil, false
}
