package service

import (
	"toDoListRestApi/src/internal/domain"
	"toDoListRestApi/src/internal/repository"
)

type TodoService interface {
	Create(todo *domain.Todo) error
	GetAllWithPagination(offset, limit int) ([]domain.Todo, error) // Изменено здесь
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) Create(todo *domain.Todo) error {

	return s.repo.Create(todo)
}

func (s *todoService) GetAllWithPagination(offset, limit int) ([]domain.Todo, error) {
	return s.repo.FindAllWithPagination(offset, limit)
}

func (s *todoService) GetByID(id uint) (*domain.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *todoService) Update(todo *domain.Todo) error {
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
