package database

import (
	"github.com/ekkapob/saucony/model"
	"github.com/golang/glog"
)

func (db *DB) Stores() (stores []model.Store) {
	query := db.Model(&stores).Order("id ASC")
	logError(query.Column("name", "phone", "city_th").Select())
	return stores
}

func (db *DB) ProductOptions() (productOptionMap map[string][]string) {
	sql := `
		SELECT DISTINCT products.style, 
		products.color
		FROM
		(SELECT collection AS style, color FROM consign
		WHERE quantity > 0
		UNION
		SELECT style, color FROM wholesales
		WHERE quantity > 0) AS products
		WHERE products.style != ''
		ORDER BY products.style
	`
	options := []struct {
		Style string
		Color string
	}{}
	_, err := db.Query(&options, sql)
	if err != nil {
		glog.Error(err)
	}
	productOptionMap = make(map[string][]string)
	for _, v := range options {
		if colors, ok := productOptionMap[v.Style]; ok {
			for _, color := range colors {
				if color != v.Color {
					productOptionMap[v.Style] = append(colors, v.Color)
				}
			}
		} else {
			productOptionMap[v.Style] = []string{v.Color}
		}
	}

	return productOptionMap
}

func (db *DB) FindStores(model string, color string, size string) (stores []model.Store) {
	sql := `
		SELECT salers.customer_no,
		stores.name,
		stores.phone,
		stores.city_th,
		concat(salers.style, ' (', salers.color, ')') AS model,
		salers.size,
		products.primary_remote_image AS remote_image,
		products.gender
		FROM stores
		LEFT JOIN
		(SELECT customer_no, size, style, color FROM consign
		UNION
		SELECT wholesaler_no, size, style, color FROM wholesales) AS salers
		ON salers.customer_no IN (stores.customer_no, stores.customer_no_2)
		LEFT JOIN products
		ON salers.style = products.model AND salers.color = products.color
		WHERE salers.style ~* ?
	`
	sql = sql + " AND salers.color ~* ? "
	if size != "" {
		sql = sql + " AND salers.size = ? "
	}
	sql = sql + " ORDER BY stores.id; "

	_, err := db.Query(&stores, sql, model, color, size)
	if err != nil {
		glog.Error(err)
	}
	return stores
}
