package checkout

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Checkout struct {
	model.Tpl
	Flash model.Flash
}

func Index(w http.ResponseWriter, r *http.Request) {
	flash := helper.GetFlash(w, r)
	t := handler.BaseTemplate("checkout.tmpl", nil)
	cart := helper.GetCart(r)
	customer := helper.GetCustomer(r)

	checkout := Checkout{
		Tpl: model.Tpl{
			Cart:     cart,
			Customer: customer,
		},
		Flash: flash,
	}
	t.ExecuteTemplate(w, "main", checkout)
}
