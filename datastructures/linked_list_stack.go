package datastructures

/*StackItem type*/
type StackItem struct {
	value interface{}
	next  *StackItem
}

/*LinkedListStack type*/
type LinkedListStack struct {
	top  *StackItem
	size int
}

/*NewLinkedListStack creates stack with specified size*/
func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{}
}

/*Len of LinkedListStack*/
func (stack *LinkedListStack) Len() int {
	return stack.size
}

/*Push new element to LinkedLinkedListStack*/
func (stack *LinkedListStack) Push(value interface{}) {
	stack.top = &StackItem{
		value: value,
		next:  stack.top,
	}

	stack.size++
}

/*Pop last element from LinkedListStack*/
func (stack *LinkedListStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	value = stack.top.value
	stack.top = stack.top.next
	stack.size--

	return value
}
