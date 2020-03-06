package datastructures

type StackItem struct {
	value interface{}
	next  *StackItem
}

type ListStack struct {
	top  *StackItem
	size int
}

func (stack *ListStack) Len() int {
	return stack.size
}

func (stack *ListStack) Push(value interface{}) {
	stack.top = &StackItem{
		value: value,
		next:  stack.top,
	}

	stack.size++
}

func (stack *ListStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	value = stack.top.value
	stack.top = stack.top.next
	stack.size--

	return value
}
