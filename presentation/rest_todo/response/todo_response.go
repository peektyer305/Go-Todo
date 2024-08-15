package rest_todo

import (
	"time"

	"github.com/google/uuid"
	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type TodoResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string             `json:"title"`
	Body        *string            `json:"body,omitempty"`
	DueDate     *time.Time         `json:"dueDate,omitempty"`
	CompletedAt *time.Time         `json:"completedAt,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func CreateTodoResponse(id valueobject.TodoId, title string, body *string, dueDate *time.Time, completedAt *time.Time, createdAt time.Time, updatedAt time.Time) TodoResponse {
	idValue,err := id.Value()
	if err != nil {
		panic(err)
	}
	return TodoResponse{
		Id:         idValue,
		Title:      title,
		Body:       body,
		DueDate:    dueDate,
		CompletedAt: completedAt,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

