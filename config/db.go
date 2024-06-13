package config

import (
	"log"

	"github.com/fazriridwan19/service-employee/domain/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "root:root@tcp(192.168.0.102:3306)/project_aka"
	log.Println("Connection to ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.Employee{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
