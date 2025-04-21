package factory

import (
	"template_app/middlewares"
	"template_app/repositories"
	"template_app/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Factory interface {
	TodoRepository() repositories.TodoRepository
	TodoService() services.TodoService
}

type factory struct {
	db *gorm.DB
}

func NewTodoFactory(ctx echo.Context) Factory {
	d := ctx.Get(middlewares.ContextMySQLKey).(*middlewares.DatabaseClient)

	return &factory{
		db: d.Session,
	}
}

func (f *factory) TodoRepository() repositories.TodoRepository {
	return repositories.NewTodoRepository(f.db)
}

func (f *factory) TodoService() services.TodoService {
	return services.NewTodoService(f.TodoRepository())
}
