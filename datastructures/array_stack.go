package datastructures

/*ListStack type*/
type ArrayStack struct {
	values []interface{}
	size   int
}

/*Length of ListStack*/
func (stack *ArrayStack) Len() int {
	return stack.size
}

/*Push new element to ListStack*/
func (stack *ArrayStack) Push(value interface{}) {
	stack.values[stack.size] = value
	stack.size++
}

/*Pop last element from ListStack*/
func (stack *ArrayStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	stack.size--
	value = stack.values[stack.size]
	stack.values[stack.size] = nil

	return value
}
