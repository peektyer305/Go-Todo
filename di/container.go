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

// providerSetに依存関係を定義
var providerSet = wire.NewSet(
	// driver
	infrastructure.NewGormPostgres,

	// Repository
	NewTodoRepository,

	// UseCase
	NewTodoUseCase,

)

func InitializeTodoUseCase() *todo.TodoUseCase {
	wire.Build(providerSet)
	return &todo.TodoUseCase{}
}

// TodoRepositoryを初期化
func NewTodoRepository(db *gorm.DB) repository.ITodoRepository {
	return &database.TodoRepository{
		Db: db,
	}
}

// TodoUseCaseを初期化
func NewTodoUseCase(todoRepo repository.ITodoRepository) *todo.TodoUseCase {
	return &todo.TodoUseCase{
		TodoRepository: todoRepo,
	}
}
