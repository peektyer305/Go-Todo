package entity

import (
	"time"
)

type TodoId string

type CreateParams struct {
    Title   string    `json:"title"`
    Body    *string   `json:"body,omitempty"`
    DueDate *time.Time `json:"dueDate,omitempty"`
}
