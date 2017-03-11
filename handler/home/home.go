package home

import (
	"html/template"
	"net/http"

	"github.com/ekkapob/saucony/database"
	"github.com/ekkapob/saucony/handler"
)

func Home(db database.DB) handler.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(handler.BaseTemplate("home.tmpl", nil))
		t.ExecuteTemplate(w, "main", nil)
	}
}
