package error

type NotFoundError struct {
	description string
}

func (ne *NotFoundError) Error() string {
	return ne.description
}

func New(description string) *NotFoundError {
	return &NotFoundError{description: description}
}

