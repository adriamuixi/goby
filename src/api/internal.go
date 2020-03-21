package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InternalAPI(r *mux.Router) {
	r.HandleFunc("/health", health)
}

func health(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, `{ "status": "OK"}`)
}