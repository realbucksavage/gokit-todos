package todos

import (
	"context"

	"github.com/jinzhu/gorm"

	"gokit-todos/database/models"
)

type Service interface {
	Get(ctx context.Context, id int) (models.Todo, error)
	GetAll(ctx context.Context) ([]models.Todo, error)
	Create(ctx context.Context, todo createRequest) (models.Todo, error)
}

type todoService struct {
	db *gorm.DB
}

func (t todoService) Get(ctx context.Context, id int) (models.Todo, error) {
	var todo models.Todo
	err := t.db.First(&todo, id).Error

	return todo, err
}

func (t todoService) GetAll(ctx context.Context) ([]models.Todo, error) {
	var todos []models.Todo
	if err := t.db.Order("id desc").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (t todoService) Create(ctx context.Context, r createRequest) (models.Todo, error) {
	todo := models.Todo{Title: r.Title, Completed: r.Completed}
	err := t.db.Create(&todo).Error

	return todo, err
}

func NewService(db *gorm.DB) Service {
	return todoService{db}
}
