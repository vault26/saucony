package database

import (
	"errors"

	"github.com/ekkapob/saucony/model"
)

func (db *DB) Promotion(code string) (promotion model.Promotion, err error) {
	query := db.Model(&promotion)
	err = query.Where("code = ?", code).Select()
	if err != nil {
		return promotion, errors.New("invalid promotion code")
	}
	return promotion, nil
}
