package repository

import (
	"gorm.io/gorm"
	"toDoListRestApi/src/internal/domain"
)

type TodoRepository interface {
	Create(todo *domain.Todo) error
	FindAllWithPagination(offset, limit int) ([]domain.Todo, error)
	FindByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindAllWithPagination(offset, limit int) ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Offset(offset).Limit(limit).Find(&todos).Error
	return todos, err
}

func (r *todoRepository) FindByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}
