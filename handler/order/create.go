package order

import (
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/handler/mail"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
	"github.com/golang/glog"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)
	session := ctx["session"].(*model.Session)
	cart := ctx["cart"].(model.Cart)

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
	model.Tmpl
	OrderRef string
}

func confirmationPage(w http.ResponseWriter, r *http.Request, orderRef string) {
	ctx := helper.GetContext(w, r)
	cart := ctx["cart"].(model.Cart)
	session := ctx["session"].(*model.Session)
	// get customer directly from session as it is not in request's context yet
	customer, _ := session.GetData("customer")
	promotion, _ := session.GetData("promotion")
	promo, _ := promotion.(model.Promotion)
	session.ClearData(w, r, "cart")

	go (func() {
		_, err := mail.OrderNotify(
			orderRef,
			cart,
			customer.(model.Customer),
			promo,
		)
		if err != nil {
			glog.Error(err)
		}
	})()

	t := handler.BaseTemplate("order_confirmation.tmpl", nil)
	orderConfirmation := OrderConfirmation{
		Tmpl: model.Tmpl{
			Cart:      cart,
			Customer:  customer.(model.Customer),
			Promotion: promo,
		},
		OrderRef: orderRef,
	}
	t.ExecuteTemplate(w, "main", orderConfirmation)
}
