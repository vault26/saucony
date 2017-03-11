package product

import (
	"net/http"

	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func Gender(w http.ResponseWriter, r *http.Request) {
	gender := mux.Vars(r)["gender"]
	if strings.ToLower(gender) == "men" {
		gender = "M"
	} else if strings.ToLower(gender) == "women" {
		gender = "W"
	}
	values := &url.Values{
		"gender[]": []string{gender},
	}
	http.Redirect(w, r, "/products?"+values.Encode(), http.StatusMovedPermanently)
}
