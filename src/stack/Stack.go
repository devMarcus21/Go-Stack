package stack

import (
	"github.com/devMarcus21/Go-Stack/src/stackErrors"
)

type Stack[T any] []T

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (stack *Stack[T]) Push(element T) bool {
	*stack = append(*stack, element)
	return true
}

func (stack *Stack[T]) Pop() (error, T) {
	if stack.IsEmpty() {
		return stackErrors.StackIsEmptyError, *new(T)
	}

	lastIndex := len(*stack) - 1             // Top index
	elementTopOfStack := (*stack)[lastIndex] // Get element top of stack

	*stack = (*stack)[:lastIndex] // Remove element from stack slice

	return nil, elementTopOfStack
}

func (stack Stack[T]) Peek() (error, T) {
	if stack.IsEmpty() {
		return stackErrors.StackIsEmptyError, *new(T)
	}

	lastIndex := len(stack) - 1
	return nil, (stack)[lastIndex]
}

func (stack Stack[T]) IsEmpty() bool {
	return len(stack) == 0
}

func (stack Stack[T]) Len() int {
	return len(stack)
}
