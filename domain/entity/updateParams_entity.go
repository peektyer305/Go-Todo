package entity

import (
	"time"
)

type UpdateParams struct {
	Title 	*string    `json:"title,omitempty"`
	Body 	*string    `json:"body,omitempty"`
	DueDate *time.Time `json:"dueDate,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`

}