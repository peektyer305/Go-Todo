package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	entity "github.com/peektyer305/Go-Todo/domain/entity"
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
	newEntity:= entity.NewTodo(t.Id,t.Title, t.Body, t.DueDate, t.CompletedAt, t.CreatedAt, t.UpdatedAt)
	fmt.Println("ModelToEntity ok")
	return newEntity
}

func FromEntityToModel(todo entity.Todo) Todo {
	 newModel:= Todo{
		Id: todo.Id,
		Title: todo.Title,
		Body: todo.Body,
		DueDate: todo.DueDate,
		CompletedAt: todo.CompletedAt,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		
	}
	fmt.Println("FromEntityToModel ok")
	return newModel
}
