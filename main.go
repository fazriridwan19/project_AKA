package main

import (
	"log"

	"github.com/fazriridwan19/service-employee/config"
	"github.com/fazriridwan19/service-employee/db"
	"github.com/fazriridwan19/service-employee/http/routes"
)

func main() {
	cfg := config.NewTest()
	db, error := db.ConnectMysql(cfg)
	if error != nil {
		log.Fatal(error)
	}
	r := routes.SetupRoutes(db)
	log.Fatal(r.Run("192.168.0.103:9090"))
}
