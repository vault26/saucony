package order

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
	db := helper.GetDB(r)
	session := helper.GetSession(r)
	cart := helper.GetCart(r)

	formValue := map[string]string{
		"firstname": r.FormValue("firstname"),
		"lastname":  r.FormValue("lastname"),
		"email":     r.FormValue("email"),
		"phone":     r.FormValue("phone"),
		"address":   r.FormValue("address"),
	}
	session.UpdateCustomerInfo(w, r, formValue)

	if len(cart.Products) < 1 {
		session.AddFlash(model.Flash{
			Type:    "danger",
			Message: "Cart is empty. กรุณาเลือกสินค้าเพื่อสั่งซื้อ",
		})
		session.Save(r, w)
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}

	errorMap, customerID, err := db.CreateCustomer(formValue)
	if err != nil {
		session.AddFlash(model.Flash{
			FormErrorMap: errorMap,
		})
		session.Save(r, w)
		http.Redirect(w, r, r.Referer()+"#customer-information", http.StatusFound)
		return
	}

	orderRef, err := db.CreateOrder(customerID, cart)
	if err != nil {
		session.AddFlash(model.Flash{
			Type:    "danger",
			Message: "Unable to place order please contact Saucony Thailand. ไม่สามารถดำเนินการสั่งซื้อได้ กรุณาติดต่อบริษัท",
		})
		session.Save(r, w)
		http.Redirect(w, r, r.Referer(), http.StatusFound)
		return
	}
	confirmationPage(w, r, orderRef)
}

type OrderConfirmation struct {
	model.Tpl
	OrderRef string
}

func confirmationPage(w http.ResponseWriter, r *http.Request, orderRef string) {
	cart := helper.GetCart(r)
	// get customer directly from session as it is not in request's context yet
	session := helper.GetSession(r)
	customer, _ := session.GetData("customer")
	session.ClearCart(w, r)

	t := handler.BaseTemplate("order_confirmation.tmpl", nil)
	orderConfirmation := OrderConfirmation{
		Tpl: model.Tpl{
			Cart:     cart,
			Customer: customer.(model.Customer),
		},
		OrderRef: orderRef,
	}
	t.ExecuteTemplate(w, "main", orderConfirmation)
}
