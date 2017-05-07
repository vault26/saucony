package handler

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/ekkapob/saucony/model"
	"github.com/leekchan/accounting"
	"github.com/russross/blackfriday"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

func init() {
	gob.Register(model.Cart{})
	gob.Register(model.Flash{})
	gob.Register(model.Customer{})
	gob.Register(model.Promotion{})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}

func BaseTemplate(contentTemplateFile string, funcMap template.FuncMap) *template.Template {
	templateFolder := "templates/dist/"
	t := template.New("base").Funcs(BaseFuncMap())
	if funcMap != nil {
		t.Funcs(funcMap)
	}
	t = template.Must(t.ParseGlob(templateFolder + "base/*.tmpl"))
	if contentTemplateFile != "" {
		t = template.Must(t.ParseFiles(fmt.Sprint(templateFolder, contentTemplateFile)))
	}
	return t
}

func BaseFuncMap() template.FuncMap {
	return template.FuncMap{
		"contains": func(list []string, text string) bool {
			for _, item := range list {
				if item == text {
					return true
				}
			}
			return false
		},
		"money": func(number float64, precision int, suffix string) string {
			ac := accounting.Accounting{Precision: precision}
			return fmt.Sprint(ac.FormatMoney(number), suffix)
		},
		"remoteImageUrl": func(key string, width interface{}) string {
			host := "http://img.wolverineworldwide.com/is/image/WolverineWorldWide/"
			host = fmt.Sprint(host, key)
			if v, ok := width.(int); ok {
				return fmt.Sprint(host, "?wid=", v)
			}
			return host
		},
		"lowerCase": func(text string) string {
			return strings.ToLower(text)
		},
		"upperCase": func(text string) string {
			return strings.ToUpper(text)
		},
		"markdown": func(text ...interface{}) template.HTML {
			output := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", text...)))
			return template.HTML(output)
		},
	}
}
