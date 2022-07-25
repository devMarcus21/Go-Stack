package stackErrors

import "errors"

var (
	StackIsEmptyError = errors.New("Stack is empty")
)
