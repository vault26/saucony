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
		(select retailer_no AS customer_no, size, style, color from retailers
		union
		select wholesaler_no, size, style, color from wholesales) AS salers
		ON salers.customer_no IN (stores.customer_no, stores.customer_no_2)
		LEFT JOIN products
		ON salers.style = products.model AND salers.color = products.color
		WHERE salers.style ~* ?
		ORDER BY stores.id;
	`
	_, err := db.Query(&stores, sql, queryText)
	if err != nil {
		glog.Error(err)
	}
	return stores
}
