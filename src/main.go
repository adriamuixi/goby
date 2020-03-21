package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"bitbucket.org/goby/api"
	"bitbucket.org/goby/config"
	"bitbucket.org/goby/logger"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	cwd, _ := os.Getwd()
	env := strings.ToLower(os.Getenv("ENVIRONMENT"))
	configPath := fmt.Sprintf("%s/settings/settings.%s.json", cwd, env)
	conf := config.LoadConfig(configPath)
	logger.InitLogger(&conf.Logger)

	router := mux.NewRouter().StrictSlash(true)
	api.InternalAPI(router)
	api.WebApi(router)

	logger.Info.Println("Starting Servers")
	err := http.ListenAndServe(":80", router)

	if err != nil {
		logger.Error.Println("Server - ERROR: ", err)
		panic(err)
	}

}
