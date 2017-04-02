package page

import (
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
)

func History(w http.ResponseWriter, r *http.Request) {
	tmpl := helper.InitTemplate(w, r)
	t := handler.BaseTemplate("history.tmpl", nil)
	t.ExecuteTemplate(w, "main", tmpl)
}
