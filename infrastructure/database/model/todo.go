package model

import (
	"time"

	"github.com/google/uuid"
	entity "github.com/peektyer305/Go-Todo/domain/entity"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type Todo struct {
	Id        uuid.UUID `gorm:"type:uuid; primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(100); not null" json:"title"`
	Body     *string `gorm:"type:text;" json:"body"`
	DueDate  *time.Time 
	CompletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) ModelToEntity() entity.Todo {
	todoId, err := valueobject.NewTodoId(t.Id.String())
	if err != nil {
		panic(err)
	}
	newEntity := entity.NewTodo(todoId, t.Title, t.Body, t.DueDate, t.CompletedAt, t.CreatedAt, t.UpdatedAt)
	return newEntity
}

func FromEntityToModel(todo entity.Todo) Todo {
	id,err := todo.Id.Value()
	if err != nil {
		panic(err)
	}
	 newModel:= Todo{
		Id: id,
		Title: todo.Title,
		Body: todo.Body,
		DueDate: todo.DueDate,
		CompletedAt: todo.CompletedAt,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		
	}
	return newModel
}
