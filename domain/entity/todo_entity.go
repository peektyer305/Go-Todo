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

func (t *Todo) GetId() valueobject.TodoId {
	return t.Id
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) GetBody() *string {
	return t.Body
}

func (t *Todo) GetDueDate() *time.Time {
	return t.DueDate
}

func (t *Todo) GetCompletedAt() *time.Time {
	return t.CompletedAt
}

func (t *Todo) GetCreatedAt() time.Time {
	return t.CreatedAt
}

func (t *Todo) GetUpdatedAt() time.Time {
	return t.UpdatedAt
}

func  NewTodo(id valueobject.TodoId, title string, body string, dueDate time.Time, completedAt time.Time, createdAt time.Time, updatedAt time.Time) Todo {
	return Todo{
		Id: id,
		Title: title,
		Body: &body,
		DueDate: &dueDate,
		CompletedAt: &completedAt,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
