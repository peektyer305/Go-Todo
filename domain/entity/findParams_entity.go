package entity

import (
	"time"

	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type FindParams struct {
	Id    *valueobject.TodoId `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
	IsCompleted *bool `json:"isCompleted,omitempty"`
	StartDate *time.Time `json:"startDate,omitempty"`
	EndDate *time.Time `json:"endDate,omitempty"`
}