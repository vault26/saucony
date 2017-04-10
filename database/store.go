package database

import (
	"github.com/ekkapob/saucony/model"
	"github.com/golang/glog"
)

func (db *DB) Stores() (stores []model.Store) {
	query := db.Model(&stores)
	logError(query.Column("name", "phone", "city_th").Select())
	return stores
}

func (db *DB) FindStores(queryText string) (stores []model.Store) {
	sql := `
		SELECT retailers.retailer_no,
		stores.name,
		stores.phone,
		stores.city_th,
		concat(retailers.style, ' (', retailers.color, ')') AS model,
		retailers.size,
		products.primary_remote_image AS remote_image
		FROM stores
		INNER JOIN retailers
		ON retailers.retailer_no IN (stores.retailer_no, stores.retailer_no_2)
		LEFT JOIN products
		ON retailers.style = products.model AND retailers.color = products.color
		WHERE retailers.style ~* ?
		GROUP BY retailers.retailer_no, stores.name, stores.phone,
		stores.city_th, retailers.style, retailers.color,
		retailers.size, stores.id, products.primary_remote_image
		ORDER BY stores.id;
	`
	_, err := db.Query(&stores, sql, queryText)
	if err != nil {
		glog.Error(err)
	}
	return stores
}
