package datastructures_test

import (
	"testing"

	"github.com/TheDoctor0/algods/datastructures"
	"github.com/stretchr/testify/assert"
)

func TestListStack(t *testing.T) {
	stack := datastructures.NewListStack()

	assert.Equal(t, 0, stack.Len(), "New stack should be empty.")

	stack.Push(1)
	assert.Equal(t, 1, stack.Len(), "Stack length should be equal 1.")

	stack.Push("test")
	assert.Equal(t, 2, stack.Len(), "Stack length should be equal 2.")

	stack.Push(3)
	assert.Equal(t, 3, stack.Len(), "Stack length should be equal 3.")

	assert.Equal(t, 3, stack.Pop(), "Value popped from stack should be equal 3.")
	assert.Equal(t, 2, stack.Len(), "Stack length should be equal 2.")

	assert.Equal(t, "test", stack.Pop(), "Value popped from stack should be equal 'test'.")
	assert.Equal(t, 1, stack.Len(), "Stack length should be equal 1.")

	assert.Equal(t, 1, stack.Pop(), "Value popped from stack should be equal 1.")
	assert.Equal(t, 0, stack.Len(), "Stack should be empty.")

	assert.Equal(t, nil, stack.Pop(), "Value popped from empty stack should be equal nil.")
}
