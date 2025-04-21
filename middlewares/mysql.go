package middlewares

import (
	"template_app/storage"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const ContextMySQLKey = "db"

type DatabaseClient struct {
	Session *gorm.DB
}

func MySQLConnectMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := storage.MySQLConnect()
			d := DatabaseClient{
				Session: session,
			}
			sqlDB, _ := d.Session.DB()

			defer sqlDB.Close()

			c.Set(ContextMySQLKey, &d)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
