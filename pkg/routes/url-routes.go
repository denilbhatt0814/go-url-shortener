package routes

import (
	"github.com/denilbhatt0814/practice/go-urls/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUrlRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/{spath}", controllers.PathHandler).Methods("GET")
	router.HandleFunc("/urls", controllers.AddPath).Methods("POST")
}
