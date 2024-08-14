package entity

import (
	"time"

	"github.com/google/uuid"
)

type FindParams struct {
	Id   * uuid.UUID `query:"id,omitempty"`
	Title *string `query:"title,omitempty"`
	Body  *string `query:"body,omitempty"`
	IsCompleted *bool `query:"isCompleted,omitempty"`
	StartDate *time.Time `query:"startDate,omitempty"`
	EndDate *time.Time `query:"endDate,omitempty"`
}