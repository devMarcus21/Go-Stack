package stackTests

import (
	"testing"

	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewStack(t *testing.T) {
	stack := stack.New[string]()

	assert.NotNil(t, stack, "Should not be nil")
	assert.Equal(t, 0, stack.Len(), "Should have no elements")
	assert.True(t, stack.IsEmpty(), "Should be empty")
}
