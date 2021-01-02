package err

import "fmt"

type TypeNotExpectedError struct {
	object interface{}
}

func (t TypeNotExpectedError) Error() string {
	return fmt.Sprintf("unexpected type %T", t.object)
}

func NewTypeNotExpectedError(object interface{}) error {
	return TypeNotExpectedError{object}
}
