package valueobject

import (
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

type TodoId struct {
	value uuid.UUID
}

func NewTodoId(value string) (TodoId, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return TodoId{},errors.New("Invalid UUID format")
	}
	return TodoId{value: id}, nil
}