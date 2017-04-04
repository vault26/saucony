package database

import "github.com/ekkapob/saucony/model"

func (db *DB) Stores() (stores []model.Store) {
	query := db.Model(&stores)
	logError(query.Select())
	return stores
}
