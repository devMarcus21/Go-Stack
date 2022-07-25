package stackTests

import (
	"testing"

	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/stretchr/testify/assert"
)

func TestShouldPushNewElementOntoStack(t *testing.T) {
	stack := stack.New[int]()
	assert.True(t, stack.IsEmpty(), "Should be empty")

	res := stack.Push(1)

	assert.True(t, res, "Should return true")
	assert.Equal(t, 1, stack.Len(), "Should have one element in stack")
	assert.False(t, stack.IsEmpty(), "Should not be empty")
}

func TestShouldPushNewElementsOntoStack(t *testing.T) {
	stack := stack.New[int]()
	assert.True(t, stack.IsEmpty(), "Should be empty")

	stack.Push(1)
	stack.Push(3)

	assert.Equal(t, 2, stack.Len(), "Should have two element in stack")
	assert.False(t, stack.IsEmpty(), "Should not be empty")
}
