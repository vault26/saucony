package page

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := helper.InitTemplate(w, r)
	t := handler.BaseTemplate("home.tmpl", nil)
	t.ExecuteTemplate(w, "main", tmpl)
}
