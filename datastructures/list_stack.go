package datastructures

/*StackItem type*/
type StackItem struct {
	value interface{}
	next  *StackItem
}

/*ListStack type*/
type ListStack struct {
	top  *StackItem
	size int
}

/*Create new ListStack*/
func NewListStack() *ListStack {
	return &ListStack{}
}

/*Get length of ListStack*/
func (stack *ListStack) Len() int {
	return stack.size
}

/*Push new element to ListStack*/
func (stack *ListStack) Push(value interface{}) {
	stack.top = &StackItem{
		value: value,
		next:  stack.top,
	}

	stack.size++
}

/*Pop last element from ListStack*/
func (stack *ListStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	value = stack.top.value
	stack.top = stack.top.next
	stack.size--

	return value
}
