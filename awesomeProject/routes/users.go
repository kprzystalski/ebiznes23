package routes

import (
	"awesomeProject/database/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Read(c echo.Context) error {
	c.Param("id")
	db, _ := c.Get("db").(*gorm.DB)
	var user models.User
	db.Where("Username = ? ", "Karol").First(&user)
	return c.JSON(http.StatusOK, user)
}
