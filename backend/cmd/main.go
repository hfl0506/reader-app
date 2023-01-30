package main

import (
	"log"
	"net/http"

	"github.com/hfl0506/reader-app/internal/config"
	"github.com/hfl0506/reader-app/internal/controller"
	"github.com/hfl0506/reader-app/internal/db"
	"github.com/hfl0506/reader-app/internal/store"
	"github.com/hfl0506/reader-app/internal/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Load config failed: %v", err)
	}

	psql, err := db.ConnectDB(c.DbUrl)

	if err != nil {
		log.Fatalf("Connect db failed: %v", err)
	}

	awsInfo := &utils.AccessPayload{
		AwsRegion:          c.AwsRegion,
		AwsAccessKeyId:     c.AwsAccessKeyId,
		AwsSecretAccessKey: c.AwsSecretAccessKey,
	}

	sess := utils.ConnectAws(awsInfo)

	e := echo.New()

	e.Validator = utils.NewValidator()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("sess", sess)
			return next(c)
		}
	})

	v1 := e.Group("/api")

	v1.GET("/healthcheck", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "health check is okay")
	})

	bs := store.NewBookStore(psql)

	us := store.NewUserStore(psql)

	h := controller.NewHandler(controller.Handler{
		UserStore: us,
		BookStore: bs,
	})

	h.RegisterRoutes(v1)

	e.Logger.Fatal(e.Start(c.Port))
}
