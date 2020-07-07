package main

import (
	"hotelAPI/config"
	"hotelAPI/main/masters"
)

func main() {
	db := config.EnvConn()
	router := config.CreateRouter()
	masters.Init(router, db)
	config.RunServer(router)
}
