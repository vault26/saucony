package handler

import (
	"html/template"
	"net/http"

	"github.com/ekkapob/saucony/database"
)

func Home(db database.DB) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(baseTemplate("home.tmpl", nil))
		t.ExecuteTemplate(w, "main", nil)
	}
}
