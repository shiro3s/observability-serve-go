package middlewares

import (
	"bytes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/trace"
)

func Init(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{},
		AllowHeaders: []string{},
	}))
	e.Use(middleware.Secure())
	e.Use(MySQLConnectMiddleware())
	e.Use(otelecho.Middleware("api-server"))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_rfc3339} [${method}]: ${status} {"uri":"${uri}","latency":"${latency}","remote_ip":"${remote_ip}","user_agent":"${user_agent}", "${custom}"}`,
		CustomTimeFormat: "2006-01-02 15:04:05",
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			span := trace.SpanFromContext(c.Request().Context())
			buf.WriteString(fmt.Sprintf("traceID: %s, spanId: %s", span.SpanContext().TraceID().String(), span.SpanContext().SpanID().String()))

			return 0, nil
		},
	}))

}
