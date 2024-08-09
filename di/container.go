//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/peektyer305/Go-Todo/application/todo"
	"github.com/peektyer305/Go-Todo/domain/repository"
	"github.com/peektyer305/Go-Todo/infrastructure"
	"github.com/peektyer305/Go-Todo/infrastructure/database"
	"gorm.io/gorm"
)

var providerSet = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// client
	//auth.NewAuthMockClient,
	// Note: ↑をコメントアウトして↓のコメントアウトを解除して wire gen すると mock2 が使われて SamplePingPong で println される文字列が変わる
	//auth.NewAuthMock2Client,

	// Repository
	NewTodoRepository,
	// queryService

	// domainService

	// useCase
	NewTodoUseCase,
)
func NewTodoRepository(db *gorm.DB) repository.ITodoRepository {
	return &database.TodoRepository{
		Db: db,
	}
}

func NewTodoUseCase(todoRepo repository.ITodoRepository) *todo.TodoUseCase {
	return &todo.TodoUseCase{
		TodoRepository: todoRepo,
	}
}
