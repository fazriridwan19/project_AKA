package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fazriridwan19/service-employee/domain/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBulk(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fileContent, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		xlsx, err := excelize.OpenReader(fileContent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var messages []string
		rows := xlsx.GetRows("Sheet1")
		for i, row := range rows {
			if i == 0 {
				continue
			}
			name := row[0]
			email := row[1]
			var employee models.Employee
			if err := db.Where("email = ?", email).First(&employee).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if employee.Id != 0 {
				message := "Email " + email + " already exist"
				messages = append(messages, message)
				continue
			}
			employee.Name = name
			employee.Email = email
			employee.CreatedAt = time.Now()
			employee.UpdatedAt = time.Now()
			if err := db.Create(&employee).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		fileContent.Close()
		if len(messages) != 0 && len(messages) != len(rows)-1 {
			c.JSON(201, gin.H{"message": "Data created with error", "error": messages})
		} else if len(messages) != 0 && len(messages) == len(rows)-1 {
			c.JSON(201, gin.H{"message": "Failed to create data", "error": messages})
		} else {
			c.JSON(201, gin.H{"message": "Data created"})
		}

	}
}
