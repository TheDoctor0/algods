package datastructures

/*ArrayStack type*/
type ArrayStack struct {
	values []interface{}
	size   int
}

/*NewArrayStack creates stack with specified size*/
func NewArrayStack(size int) *ArrayStack {
	return &ArrayStack{values: make([]interface{}, size)}
}

/*Len of ArrayStack (current)*/
func (stack *ArrayStack) Len() int {
	return stack.size
}

/*Push new element to ArrayStack*/
func (stack *ArrayStack) Push(value interface{}) {
	if stack.Len() == len(stack.values) {
		panic("stack size exceeded")
	}

	stack.values[stack.size] = value
	stack.size++
}

/*Pop last element from ArrayStack*/
func (stack *ArrayStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	stack.size--
	value = stack.values[stack.size]
	stack.values[stack.size] = nil

	return value
}
