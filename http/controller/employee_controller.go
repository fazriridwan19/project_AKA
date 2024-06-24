package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fazriridwan19/service-employee/domain/models"
	"github.com/fazriridwan19/service-employee/domain/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CreateBulk(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeRepo := repositories.NewEmployeeRepository(db)
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
			employee := models.Employee{
				Name:  name,
				Email: email,
			}
			err := employeeRepo.Create(c, &employee)
			if employee.Id != 0 {
				message := "Email " + email + " already exist"
				messages = append(messages, message)
				continue
			}
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
