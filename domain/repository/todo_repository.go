package repository

import (
	"context"

	"github.com/peektyer305/Go-Todo/domain/entity"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type ITodoRepository interface {
	FindById(ctx context.Context, id valueobject.TodoId) (*entity.Todo, error)
	FindAllByQuery(ctx context.Context, queries entity.FindParams) ([]entity.Todo, error)
	//Save(ctx context.Context, todo entity.Todo) (*entity.Todo, error)
	//DeleteById(ctx context.Context, id valueobject.TodoId) error

}
