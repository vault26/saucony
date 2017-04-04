package page

import (
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
	"github.com/ekkapob/saucony/helper"
	"github.com/ekkapob/saucony/model"
)

type Stores struct {
	model.Tmpl
	Stores []model.Store
}

func Store(w http.ResponseWriter, r *http.Request) {
	ctx := helper.GetContext(w, r)
	db := ctx["db"].(database.DB)

	// stores := db.Stores()
	// fmt.Println(stores)

	tmpl := helper.InitTemplate(w, r)
	stores := &Stores{
		Tmpl:   tmpl,
		Stores: db.Stores(),
	}

	t := handler.BaseTemplate("store.tmpl", nil)
	t.ExecuteTemplate(w, "main", stores)
}
