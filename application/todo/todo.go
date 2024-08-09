package todo

import (
	"context"

	"github.com/peektyer305/Go-Todo/domain/entity"
	"github.com/peektyer305/Go-Todo/domain/repository"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type TodoUseCase struct {
	TodoRepository *repository.ITodoRepository
}

func (t *TodoUseCase) FindAllByCriterias(ctx context.Context, f entity.FindParams) ([]entity.Todo, error){
	todos,err := t.TodoRepository.FindAllByQuery(f)	
	if err != nil {
		return nil, err
	}
	 for _, todo := range todos {
		todo.
}
}

func (t *TodoUseCase) FindById(ctx context.Context id valueobject.TodoId) (*entity.Todo, error) {
	todo,err:= t.TodoRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
