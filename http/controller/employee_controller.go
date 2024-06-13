package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(201, gin.H{"data": "Created"})
	}
}
