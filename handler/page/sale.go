package page

import "net/http"

func Sale(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/products?gender%5B%5D=M&gender%5B%5D=W&section%5B%5D=sale", http.StatusMovedPermanently)
}
