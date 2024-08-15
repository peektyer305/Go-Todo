package todo

import (
	"context"

	"github.com/google/uuid"
	"github.com/peektyer305/Go-Todo/domain/entity"
	"github.com/peektyer305/Go-Todo/domain/repository"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
	request "github.com/peektyer305/Go-Todo/presentation/rest_todo/request"
)

type TodoUseCase struct {
	TodoRepository repository.ITodoRepository
}

func (t *TodoUseCase) FindAllByCriterias(ctx context.Context, f request.FindParams) ([]entity.Todo, error){
	todos,err := t.TodoRepository.FindAllByQuery(ctx,f)	
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *TodoUseCase) FindById(ctx context.Context, id valueobject.TodoId) (*entity.Todo, error) {
	todo,err:= t.TodoRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (t *TodoUseCase) Create (ctx context.Context, params request.CreateParams) (*entity.Todo, error) {
	
	todoId, err := valueobject.NewTodoId(uuid.New().String())
	if err != nil {
		return nil, err
	}
	
	todo := entity.Todo{
		Id:      todoId,
		Title:   params.Title,
		Body:    params.Body,
		DueDate: params.DueDate,
	}
	createdTodo, err := t.TodoRepository.Save(ctx, todo)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}

func (t *TodoUseCase) UpdateById (ctx context.Context, id valueobject.TodoId, params request.UpdateParams) (*entity.Todo, error) {
	todo, err := t.TodoRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if params.Title != nil {
		todo.Title = *params.Title
	}
	if params.Body != nil {
		todo.Body = params.Body
	}
	if params.DueDate != nil {
		todo.DueDate = params.DueDate
	}
	if params.CompletedAt != nil {
		todo.CompletedAt = params.CompletedAt
	}
	updatedTodo, err := t.TodoRepository.Save(ctx, *todo)
	if err != nil {
		return nil, err
	}
	return updatedTodo, nil
}

func (t *TodoUseCase) DeleteById (ctx context.Context,id valueobject.TodoId) error {
	err := t.TodoRepository.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (t *TodoUseCase) CopyById (ctx context.Context, id valueobject.TodoId) (*entity.Todo, error) {
	targetTodo, err := t.TodoRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	newTodoId, err := valueobject.NewTodoId(uuid.New().String())
		if err != nil {
			return nil, err
		}
	copiedTodo := entity.Todo{
		Id:    newTodoId,
		Title: targetTodo.Title + "(のコピー)",
	}
	createdTodo, err := t.TodoRepository.Save(ctx, copiedTodo)
	if err != nil {
		return nil, err
	}
	return createdTodo, nil
}