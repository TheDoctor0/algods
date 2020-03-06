package datastructures

/*ArrayStack type*/
type ArrayStack struct {
	values  []interface{}
	current int
	size    int
}

/*NewArrayStack creates stack with specified size*/
func NewArrayStack(size int) *ArrayStack {
	return &ArrayStack{size: size, values: make([]interface{}, size)}
}

/*Len of ArrayStack (current)*/
func (stack *ArrayStack) Len() int {
	return stack.current
}

/*Push new element to ArrayStack*/
func (stack *ArrayStack) Push(value interface{}) {
	stack.values[stack.current] = value
	stack.current++
}

/*Pop last element from ArrayStack*/
func (stack *ArrayStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	stack.current--
	value = stack.values[stack.current]
	stack.values[stack.current] = nil

	return value
}
