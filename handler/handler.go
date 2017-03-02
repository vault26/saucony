package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type handleFunc func(http.ResponseWriter, *http.Request)

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}

func baseTemplate(contentTemplateFile string, funcMap template.FuncMap) (*template.Template, error) {
	t := template.Must(template.ParseGlob("templates/base/*.tmpl"))
	if funcMap != nil {
		t = t.Funcs(funcMap)
	}
	return t.ParseFiles(fmt.Sprintf("templates/%v", contentTemplateFile))
}
