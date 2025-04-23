package main

import (
	"time"

	"github.com/labstack/echo/v4"

	"template_app/middlewares"
	"template_app/routes"
	"template_app/utils"
)

func main() {
	cleanup, err := utils.SetupTraceProvider(10 * time.Second)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	e := echo.New()
	middlewares.Init(e)
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
