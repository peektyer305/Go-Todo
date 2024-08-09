package entity

import (
	"time"

	valueobject "github.com/peektyer305/Go-Todo/domain/value_object"
)

type FindParams struct {
	Id    *valueobject.TodoId
	Title *string
	Body  *string
	IsCompleted *bool
	StartDate *time.Time
	EndDate *time.Time
}