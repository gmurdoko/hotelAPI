package config

import (
	"fmt"
	"hotelAPI/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateRouter creating new router for app.goCreateRouter
func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

//RunServer running the server
func RunServer(router *mux.Router) {
	appHost := utils.ViperGetEnv("APP_HOST", "urhost")
	appPort := utils.ViperGetEnv("APP_PORT", "urport")
	hostListen := fmt.Sprintf("%v:%v", appHost, appPort)
	log.Printf("Ready to listen on %v", hostListen)
	err := http.ListenAndServe(hostListen, router)
	if err != nil {
		log.Fatal(err)
	}
}
