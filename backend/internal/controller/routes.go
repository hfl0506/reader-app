package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	Db *gorm.DB
}

func (c *Controller) RegisterRoutes(db *gorm.DB, e *echo.Echo) {
	appRoutes := e.Group("/api")

	bookRoutes := appRoutes.Group("/books")
	bookRoutes.GET("/", c.GetBook)
}