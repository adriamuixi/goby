package api

import (
	"bitbucket.org/goby/controllers"
	"github.com/gorilla/mux"
)

func WebApi(r *mux.Router) {
	brandRouter := r.PathPrefix("/es").Subrouter()

	brandRouter.HandleFunc("/home", controllers.Home).Methods("GET")
}
