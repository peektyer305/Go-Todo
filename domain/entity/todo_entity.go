package entity

import (
	"time"

	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type Todo struct {
	Id    valueobject.TodoId
	Title      string
	Body 	 *string
	DueDate   *time.Time
	CompletedAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func  NewTodo(id valueobject.TodoId, title string, body *string, dueDate *time.Time, completedAt *time.Time, createdAt time.Time, updatedAt time.Time) Todo {
	return Todo{
		Id: id,
		Title: title,
		Body: body,
		DueDate: dueDate,
		CompletedAt: completedAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}


		
