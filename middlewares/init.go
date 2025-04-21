package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_rfc3339} [${method}]: ${status} {"uri":"${uri}","latency":"${latency}","remote_ip":"${remote_ip}","user_agent":"${user_agent}"}`,
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{},
		AllowHeaders: []string{},
	}))
	e.Use(middleware.Secure())
	e.Use(MySQLConnectMiddleware())
}
