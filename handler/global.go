package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(baseTemplate("home.tmpl"))
	t.ExecuteTemplate(w, "main", nil)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}

func baseTemplate(contentTmpl string) (*template.Template, error) {
	return template.ParseFiles(
		"templates/header.tmpl",
		"templates/main.tmpl",
		"templates/footer.tmpl",
		"templates/scripts.tmpl",
		"templates/styles.tmpl",
		fmt.Sprintf("templates/%v", contentTmpl),
	)
}
