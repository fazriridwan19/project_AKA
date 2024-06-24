package routes

import (
	"github.com/fazriridwan19/service-employee/http/controller"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(db *sqlx.DB) *gin.Engine {
	routes := gin.Default()

	routes.POST("/employee/create/bulk", controller.CreateBulk(db))

	return routes
}
