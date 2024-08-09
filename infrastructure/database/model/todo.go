package model

import (
	"time"

	entity "github.com/peektyer305/Go-Todo/domain/entity"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type Todo struct {
	Id        valueobject.TodoId `gorm:"type:uuid; primaryKey" json:"id"`
	Title    string `gorm:"type:varchar(100); not null" json:"title"`
	Body     *string `gorm:"type:text;" json:"body"`
	DueDate  *time.Time 
	CompletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Todo) ToEntity() entity.Todo {
	return entity.NewTodo( t.Id, t.Title, *t.Body, *t.DueDate, *t.CompletedAt, t.CreatedAt, t.UpdatedAt)
}

func NewTodoFromEntity(todo entity.Todo) Todo {
	return Todo{
		Id: todo.Id(),
		Title: todo.Title(),
		Body: todo.Body(),
		DueDate: todo.DueDate(),
		CompletedAt: todo.CompletedAt(),
		CreatedAt: todo.CreatedAt(),
		UpdatedAt: todo.UpdatedAt(),
		
	}
}
