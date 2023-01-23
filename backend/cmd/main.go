package main

import (
	"log"

	"github.com/hfl0506/reader-app/internal/config"
	"github.com/hfl0506/reader-app/internal/controller"
	"github.com/hfl0506/reader-app/internal/db"
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

	e := echo.New()

	controller := &controller.Controller{
		Db: psql,
	}

	e.Logger.Fatal(e.Start(":1323"))
}