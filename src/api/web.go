package api

import (
	"github.com/adriamuixi/goby/controllers"
	"github.com/gorilla/mux"
)

func WebApi(r *mux.Router) {
	brandRouter := r.PathPrefix("/es").Subrouter()

	brandRouter.HandleFunc("/home", controllers.Home).Methods("GET")
}
