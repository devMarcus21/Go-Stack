package stackTests

import (
	"testing"

	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/devMarcus21/Go-Stack/src/stackErrors"
	"github.com/stretchr/testify/assert"
)

func TestShouldPeekFromStack_StackEmpty(t *testing.T) {
	stack := stack.New[int]()
	assert.True(t, stack.IsEmpty(), "Should be empty")

	err, res := stack.Peek()

	assert.Equal(t, 0, res, "Should be default value")
	assert.NotNil(t, err, "Should return non nil error")
	assert.Equal(t, stackErrors.StackIsEmptyError, err, "Should return an empty stack error")
}

func TestShouldPeekFromStack(t *testing.T) {
	stack := stack.New[int]()
	stack.Push(2)
	assert.False(t, stack.IsEmpty(), "Should not be empty")

	err, res := stack.Peek()

	assert.Equal(t, 2, res, "Should be top of stack value")
	assert.Nil(t, err, "Should return nil error")
}
