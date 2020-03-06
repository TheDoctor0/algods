package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := NewListStack()

	assert.Equal(t, stack.Len(), 0, "New stack should be empty.")

	stack.Push(1)
	assert.Equal(t, stack.Len(), 1, "Stack length should be equal 1.")

	stack.Push("test")
	assert.Equal(t, stack.Len(), 2, "Stack length should be equal 2.")

	stack.Push(3)
	assert.Equal(t, stack.Len(), 3, "Stack length should be equal 3.")

	assert.Equal(t, stack.Pop(), 3, "Value popped from stack should be equal 3.")
	assert.Equal(t, stack.Len(), 2, "Stack length should be equal 2.")

	assert.Equal(t, stack.Pop(), "test", "Value popped from stack should be equal 'test'.")
	assert.Equal(t, stack.Len(), 1, "Stack length should be equal 1.")

	assert.Equal(t, stack.Pop(), 1, "Value popped from stack should be equal 1.")
	assert.Equal(t, stack.Len(), 0, "Stack should be empty.")
}
