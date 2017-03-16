package page

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	db := helper.GetDB(r)
	_ = db
	cart := helper.GetCart(r)
	t := handler.BaseTemplate("home.tmpl", nil)
	t.ExecuteTemplate(w, "main", model.Tpl{Cart: cart})
}
