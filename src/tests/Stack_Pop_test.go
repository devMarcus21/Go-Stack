package stackTests

import (
	"testing"

	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/devMarcus21/Go-Stack/src/stackErrors"
	"github.com/stretchr/testify/assert"
)

func TestShouldNotPopFromStack_StackEmpty(t *testing.T) {
	stack := stack.New[int]()
	assert.True(t, stack.IsEmpty(), "Should be empty")

	err, res := stack.Pop()

	assert.Equal(t, 0, res, "Should be default value")
	assert.NotNil(t, err, "Should return non nil error")
	assert.Equal(t, stackErrors.StackIsEmptyError, err, "Should return an empty stack error")
}

func TestShouldPopFromStack(t *testing.T) {
	stack := stack.New[int]()
	assert.True(t, stack.IsEmpty(), "Should be empty")
	stack.Push(3)
	assert.Equal(t, 1, stack.Len(), "Should have element in stack")

	err, res := stack.Pop()

	assert.Nil(t, err, "Should not return an error")
	assert.Equal(t, 3, res, "Should be element at top of stack")
	assert.True(t, stack.IsEmpty(), "Should be now empty")

	err, res = stack.Pop()

	assert.Equal(t, 0, res, "Should now be default value")
	assert.NotNil(t, err, "Should now return non nil error")
	assert.Equal(t, stackErrors.StackIsEmptyError, err, "Should now return an empty stack error")
}
