package database

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ekkapob/saucony/model"
	"github.com/golang/glog"
	pg "gopkg.in/pg.v5"
	"gopkg.in/pg.v5/orm"
)

func (db *DB) Products(params map[string][]string) (products []model.Product) {
	query := db.Model(&products)
	if queryText, ok := paramsValue(params, "query"); ok {
		setProductSearchQuery(query, queryText[0])
	}
	if genders, ok := paramsValue(params, "genders"); ok {
		query.Where("gender IN (?)", pg.In(genders))
	}
	if sizes, ok := paramsValue(params, "sizes"); ok {
		query.Where("sizes && ?", pg.Array(sizes))
	}
	if types, ok := paramsValue(params, "types"); ok {
		query.Where("types && ?", pg.Array(types))
	}
	if features, ok := paramsValue(params, "features"); ok {
		for k, v := range features {
			if v == "sale" {
				query.Where("discount > 0")
				features = append(features[:k], features[k+1:]...)
			}
		}
		if len(features) > 0 {
			query.Where("features && ?", pg.Array(features))
		}
	}
	logError(query.Order("model").Select())
	products = db.filterProducts(products)
	return products
}

func (db *DB) avaiableProductIdMap() (idMap map[int]int) {
	var ids []int
	sql := `
		SELECT products.id
		FROM products
		INNER JOIN consign 
		ON consign.style = products.model
		WHERE	customer_no IN ('11112', '11111')
		AND consign.quantity > 0
		GROUP BY products.id;
	`
	_, err := db.Query(&ids, sql)
	if err != nil {
		glog.Error(err)
		return idMap
	}
	idMap = make(map[int]int)
	for _, v := range ids {
		idMap[v] = v
	}
	return idMap
}

func (db *DB) availableProductSizes(product model.Product) (sizes []string) {
	sql := `
		SELECT size
		FROM consign
		WHERE	customer_no IN ('11112', '11111')
		AND style = ?
		AND color = ?
		AND gender = ?
		AND quantity > 0
	`
	_, err := db.Query(&sizes, sql, product.Model, product.Color, product.Gender)
	if err != nil {
		glog.Error(err)
	}
	return sizes
}

// Filter for stock available products
func (db *DB) filterProducts(products []model.Product) []model.Product {
	avaiableProductIdMap := db.avaiableProductIdMap()
	for k, v := range products {
		if _, ok := avaiableProductIdMap[v.ID]; !ok {
			products = append(products[:k], products[k+1:]...)
		}
	}
	return products
}

func (db *DB) ModelProducts(params map[string]string) (products []model.Product) {
	query := db.Model(&products)
	if params["color"] != "" {
		query.Where("color = ?", params["color"])
	}
	if params["gender"] != "" {
		query.Where("gender = ?", params["gender"])
	}
	logError(query.Where("model_path = ?", params["model_path"]).Select())
	for k, v := range products {
		products[k].Sizes = db.availableProductSizes(v)
	}
	return products
}

func (db *DB) Product(params map[string]string) (product *model.Product, err error) {
	if params["id"] != "" {
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			return nil, errors.New("no ID provided")
		}
		product = &model.Product{ID: id}
		if err := db.Select(&product); err != nil {
			return nil, errors.New("not found")
		}
	}
	return product, nil
}

func setProductSearchQuery(query *orm.Query, queryText string) {
	queryText = prepareSearchQuery(queryText)
	query.Where("color ~* (?)", queryText)
	query.WhereOr("model ~* (?)", queryText)
}

// Output to pattern (A|B|C) from "A B C"
func prepareSearchQuery(queryText string) string {
	reg := regexp.MustCompile(`/s+`)
	queryText = reg.ReplaceAllString(queryText, " ")
	queryText = strings.Trim(queryText, " ")
	queryText = strings.Replace(queryText, " ", "|", -1)
	return fmt.Sprint("(", queryText, ")")
}

func paramsValue(params map[string][]string, key string) ([]string, bool) {
	if value := params[key]; len(value) > 0 {
		return value, true
	}
	return nil, false
}
