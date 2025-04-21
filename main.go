package main

import (
	"github.com/labstack/echo"

	"template_app/middlewares"
	"template_app/routes"
)

func main() {
	e := echo.New()
	middlewares.Init(e)
	routes.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
