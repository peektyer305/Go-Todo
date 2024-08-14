package rest_todo

import (
	"time"

	"github.com/google/uuid"
)

type TodoResponse struct {
	Id         uuid.UUID `json:"id"`
	Title       string             `json:"title"`
	Body        *string            `json:"body,omitempty"`
	DueDate     *time.Time         `json:"dueDate,omitempty"`
	CompletedAt *time.Time         `json:"completedAt,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

func CreateTodoResponse(id uuid.UUID, title string, body *string, dueDate *time.Time, completedAt *time.Time, createdAt time.Time, updatedAt time.Time) TodoResponse {
	return TodoResponse{
		Id:         id,
		Title:      title,
		Body:       body,
		DueDate:    dueDate,
		CompletedAt: completedAt,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

