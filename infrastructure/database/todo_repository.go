package database

import (
	"context"

	"github.com/peektyer305/Go-Todo/domain/entity"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
	"github.com/peektyer305/Go-Todo/infrastructure/database/model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)


type TodoRepository struct {
	Db *gorm.DB
}

func (t *TodoRepository) FindById (ctx context.Context, id valueobject.TodoId) (*entity.Todo, error) {
	conn  := t.Db.WithContext(ctx)
	var todoModel model.Todo
	if err := conn.Where("id = ?", id).First(&todoModel).Error; err != nil {
		return nil, err
	}
	return lo.ToPtr(todoModel.ToEntity()), nil
} 

func (t *TodoRepository) FindAllByQuery(ctx context.Context, params entity.FindParams) ([]entity.Todo, error) {
	conn := t.Db.WithContext(ctx)
	var todoModels []model.Todo

	query := conn.Model(&model.Todo{})

	if params.Id != nil {
		query = query.Where("id = ?", params.Id)
	}

	if params.Title != nil {
		query = query.Where("title = ?", "%"+*params.Title+"%")
	}

	if params.Body != nil {
		query = query.Where("body = ?",  "%"+*params.Body+ "%")
	}

	if params.IsCompleted != nil {
		query = query.Where("conpleted_at IS", !*params.IsCompleted)
	}

	if params.StartDate != nil {
		query = query.Where("created_at >= ?", params.StartDate)
	}

	if params.EndDate != nil {
		query = query.Where("created_at <= ?", params.EndDate)
	}

	if err := query.Find(&todoModels).Error; err != nil {
		return nil, err
	}

	var todoEntities []entity.Todo
	for _, todoModel := range todoModels {
		todoEntities = append(todoEntities, todoModel.ToEntity())
	}

	return todoEntities, nil
}

func (t *TodoRepository) Save(ctx context.Context, todo entity.Todo) (*entity.Todo, error) {
	conn := t.Db.WithContext(ctx)
	todoModel := model.NewTodoFromEntity(todo)
	if err := conn.Save(&todoModel).Error; err != nil {
		return nil, err
	}
	return lo.ToPtr(todoModel.ToEntity()), nil
}

func (t *TodoRepository) DeleteById(ctx context.Context, id valueobject.TodoId) error {
	conn := t.Db.WithContext(ctx)
	if err := conn.Delete(&model.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

