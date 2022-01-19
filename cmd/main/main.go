package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/denilbhatt0814/practice/go-urls/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server...")
	r := mux.NewRouter()
	routes.RegisterUrlRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("192.168.29.98:8989", r))
}
