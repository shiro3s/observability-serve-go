package repositories

import (
	"template_app/dao"
	"template_app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll(cond models.TodoSearchParam) ([]models.Todo, int64)
	FindById(id int, cond models.TodoParam) (*models.Todo, error)
	Create(postBody models.TodoBody) error
	Update(id int, putBody models.TodoBody) error
	Delete(id int) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db,
	}
}

func (r *todoRepository) FindAll(cond models.TodoSearchParam) ([]models.Todo, int64) {
	todos, count := dao.SearchTodo(r.db, cond)
	return todos, count
}

func (r *todoRepository) FindById(id int, cond models.TodoParam) (*models.Todo, error) {
	return dao.FindTodoById(r.db, id, cond)
}

func (r *todoRepository) Create(postBody models.TodoBody) error {
	return dao.CreateTodo(r.db, postBody)
}

func (r *todoRepository) Update(id int, putBody models.TodoBody) error {
	return dao.UpdateTodo(r.db, id, putBody)
}

func (r *todoRepository) Delete(id int) error {
	return dao.DeleteTodo(r.db, id)
}
