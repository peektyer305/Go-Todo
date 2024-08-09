package entity

import (
	"time"

	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type Todo struct {
	id    valueobject.TodoId    
	title      string
	body 	 *string
	dueDate   *time.Time
	completedAt *time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (t *Todo) Id() valueobject.TodoId {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Body() *string {
	return t.body
}

func (t *Todo) DueDate() *time.Time {
	return t.completedAt
}

func (t *Todo) CompletedAt() *time.Time {
	return t.completedAt
}

func (t *Todo) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Todo) UpdatedAt() time.Time {
	return t.updatedAt
}

func  NewTodo(id valueobject.TodoId, title string, body string, dueDate time.Time, completedAt time.Time, createdAt time.Time, updatedAt time.Time) Todo {
	return Todo{
		id: id,
		title: title,
		body: &body,
		dueDate: &dueDate,
		completedAt: &completedAt,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
