package api

import (
	"github.com/goby/controllers"
	"github.com/gorilla/mux"
)

func WebApi(r *mux.Router) {
	brandRouter := r.PathPrefix("/es").Subrouter()

	brandRouter.HandleFunc("/home", controllers.Home).Methods("GET")
}
