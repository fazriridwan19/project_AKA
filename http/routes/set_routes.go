package routes

import (
	"github.com/fazriridwan19/service-employee/http/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	routes := gin.Default()

	routes.POST("/employee/create/bulk", controller.Create(db))

	return routes
}
