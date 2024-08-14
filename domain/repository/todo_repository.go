package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/peektyer305/Go-Todo/domain/entity"
	request "github.com/peektyer305/Go-Todo/presentation/rest_todo/request"
)

type ITodoRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*entity.Todo, error)
	FindAllByQuery(ctx context.Context, queries request.FindParams) ([]entity.Todo, error)
	Save(ctx context.Context, todo entity.Todo) (*entity.Todo, error)
	DeleteById(ctx context.Context, id uuid.UUID) error
	
}
