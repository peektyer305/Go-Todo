package valueobject

import "github.com/google/uuid"

type TodoId struct {
	value uuid.UUID
}

func NewTodoId(value uuid.UUID) TodoId {
	return TodoId{
		value: value,
	}
}