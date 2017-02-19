package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ekkapob/saucony/handler"
	"github.com/gorilla/mux"
	"github.com/namsral/flag"
)

func main() {
	port := flag.Int("port", 8080, "port")
	flag.Parse()

	r := mux.NewRouter()
	r.PathPrefix("/assets/").Handler(
		http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))),
	)
	r.HandleFunc("/", handler.Home)
	r.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	portname := fmt.Sprintf(":%v", *port)
	fmt.Println("Server is running on", portname)
	log.Fatal(http.ListenAndServe(portname, r))
}
