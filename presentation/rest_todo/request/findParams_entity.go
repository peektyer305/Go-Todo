package entity

import (
	"time"

	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type FindParams struct {
	Id   * valueobject.TodoId `query:"id,omitempty"`
	Title *string `query:"title,omitempty"`
	Body  *string `query:"body,omitempty"`
	IsCompleted *bool `query:"isCompleted,omitempty"`
	StartDate *time.Time `query:"startDate,omitempty"`
	EndDate *time.Time `query:"endDate,omitempty"`
}