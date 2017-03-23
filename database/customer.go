package database

import (
	"github.com/asaskevich/govalidator"
	"github.com/ekkapob/saucony/model"
)

func (db *DB) CreateCustomer(
	params map[string]string) (errorMap map[string]string, id int, err error) {
	customer := &model.Customer{
		Firstname: params["firstname"],
		Lastname:  params["lastname"],
		Email:     params["email"],
		Phone:     params["phone"],
		Address:   params["address"],
	}

	_, err = govalidator.ValidateStruct(customer)
	if err != nil {
		errorMap = govalidator.ErrorsByField(err)
		return
	}
	err = db.Insert(customer)
	if err != nil {
		return
	}
	return errorMap, customer.ID, nil
}
