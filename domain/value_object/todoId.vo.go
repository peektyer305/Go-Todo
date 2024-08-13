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

func (t *TodoId) Scan(value interface{}) error {
	if val, ok := value.(string); ok {
		id, err := uuid.Parse(val)
		if err != nil {
			return errors.New("Invalid UUID format")
		}
		t.value = id
		return nil
	}
	return errors.New("failed to scan TodoId")
}

func (t TodoId) Value() (interface{}, error) {
	return t.value.String(), nil
}

func (t TodoId) String() string {
	return t.value.String()
}