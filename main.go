package main

import (
	"log"

	"github.com/fazriridwan19/service-employee/config"
	"github.com/fazriridwan19/service-employee/http/routes"
)

func main() {
	db := config.ConnectDatabase()
	r := routes.SetupRoutes(db)
	log.Fatal(r.Run("192.168.0.101:9090"))
}
