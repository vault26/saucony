package checkout

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cart := helper.GetCart(r)
	t := handler.BaseTemplate("checkout.tmpl", nil)
	t.ExecuteTemplate(w, "main", model.Tpl{Cart: cart})
}
