package handler

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "123test")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "not found")
}
