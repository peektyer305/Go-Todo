package database

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/peektyer305/Go-Todo/domain/entity"
	myError "github.com/peektyer305/Go-Todo/errors"
	"github.com/peektyer305/Go-Todo/infrastructure/database/model"
	request "github.com/peektyer305/Go-Todo/presentation/rest_todo/request"
	"gorm.io/gorm"
)


type TodoRepository struct {
	Db *gorm.DB
}

func (t *TodoRepository) FindById (ctx context.Context, id uuid.UUID) (*entity.Todo, error) {
	conn := t.Db.WithContext(ctx)
	var todoModel model.Todo
	fmt.Println("FindById")
	if err := conn.Where("id = ?", id.String()).First(&todoModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
            return nil, myError.New("todo not found")
        }
		return nil, err
		}
	fmt.Println("FindById ok")
	entity := todoModel.ModelToEntity()
	return &entity, nil
} 

func (t *TodoRepository) FindAllByQuery(ctx context.Context, params request.FindParams) ([]entity.Todo, error) {
	conn := t.Db.WithContext(ctx)
	var todoModels []model.Todo

	query := conn.Model(&model.Todo{})

	if params.Id != nil {
		query = query.Where("id = ?", *params.Id)
	}

	if params.Title != nil {
		query = query.Where("title LIKE ?", "%"+*params.Title+"%")
	}

	if params.Body != nil {
		query = query.Where("body LIKE ?",  "%"+*params.Body+ "%")
	}
	
	if params.IsCompleted != nil {
		if *params.IsCompleted {
			query = query.Where("completed_at IS NOT NULL")
		} else {
			query = query.Where("completed_at IS NULL")
		}
	}	

	if params.StartDate != nil {
		query = query.Where("created_at >= ?", *params.StartDate)
	}

	if params.EndDate != nil {
		query = query.Where("created_at <= ?", *params.EndDate)
	}

	if err := query.Find(&todoModels).Error; err != nil {
		return nil, err
	}
	fmt.Println("query",query)
	var todoEntities []entity.Todo
	fmt.Println("todos:",todoModels)
	for _, todoModel := range todoModels {
		todoEntities = append(todoEntities, todoModel.ModelToEntity())
	}

	return todoEntities, nil
}

func (t *TodoRepository) Save(ctx context.Context, todo entity.Todo) (*entity.Todo, error) {
	conn := t.Db.WithContext(ctx)
	todoModel := model.FromEntityToModel(todo)
	if err := conn.Save(&todoModel).Error; err != nil {
		return nil, err
	}
	fmt.Println("Save ok")
	return &todo, nil
}

func (t *TodoRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	conn := t.Db.WithContext(ctx)
	if err := conn.Delete(&model.Todo{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return myError.New("todo not found")
		}
		return err
	}
	return nil
}

