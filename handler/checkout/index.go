package checkout

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Checkout struct {
	model.Tmpl
	Flash model.Flash
}

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	flash := ctx["flash"].(model.Flash)
	cart := ctx["cart"].(model.Cart)
	customer := ctx["customer"].(model.Customer)
	t := handler.BaseTemplate("checkout.tmpl", nil)

	checkout := Checkout{
		Tmpl: model.Tmpl{
			Cart:     cart,
			Customer: customer,
		},
		Flash: flash,
	}
	t.ExecuteTemplate(w, "main", checkout)
}
