package middlewares

import (
	"fmt"
	"template_app/utils"

	"github.com/grafana/loki-client-go/loki"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/common/model"

	"go.opentelemetry.io/otel/trace"
)

func LokiLoggerMiddleware(client *loki.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			spanCtx := trace.SpanContextFromContext(ctx)

			fmt.Println(c.Request())

			labels := model.LabelSet{
				"service_name": "api-service",
				"path":         model.LabelValue(c.Request().URL.Path),
				"method":       model.LabelValue(c.Request().Method),
				"trace_id":     model.LabelValue(spanCtx.SpanID().String()),
				"query":        model.LabelValue(c.Request().URL.Query().Encode()),
			}

			_ = utils.SendLoki(client, labels)
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
