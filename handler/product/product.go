package product

import (
	"fmt"
	"html/template"

	"github.com/leekchan/accounting"
)

func productsFuncMap() template.FuncMap {
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
	}
}

// func Etag(resource string, w http.ResponseWriter, r *http.Request) bool {
// 	// TODO: generate dynamic etag for dynamic content
// 	key := fmt.Sprint(`"`, resource, ":", r.URL.RawQuery, `"`)
// 	w.Header().Set("Etag", key)
// 	w.Header().Set("Cache-Control", "max-age=86400") // 1 hour (60*60*24)
// 	if match := r.Header.Get("If-None-Match"); match != "" {
// 		if strings.Contains(match, key) {
// 			w.WriteHeader(http.StatusNotModified)
// 			return true
// 		}
// 	}
// 	return false
// }
