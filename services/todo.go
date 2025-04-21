package services

import (
	"template_app/models"
	"template_app/repositories"
)

type TodoService interface {
	FindAll(cond models.TodoSearchParam) ([]models.Todo, int64)
	FindById(id int, cond models.TodoParam) (*models.Todo, error)
	Create(postBody models.TodoBody) error
	Update(id int, putBody models.TodoBody) error
	Delete(id int) error
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &todoService{
		todoRepository,
	}
}

func (s *todoService) FindAll(cond models.TodoSearchParam) ([]models.Todo, int64) {
	return s.todoRepository.FindAll(cond)
}

func (s *todoService) FindById(id int, cond models.TodoParam) (*models.Todo, error) {
	return s.todoRepository.FindById(id, cond)
}

func (s *todoService) Create(postBody models.TodoBody) error {
	return s.todoRepository.Create(postBody)
}

func (s *todoService) Update(id int, putBody models.TodoBody) error {
	return s.todoRepository.Update(id, putBody)
}

func (s *todoService) Delete(id int) error {
	return s.todoRepository.Delete(id)
}
