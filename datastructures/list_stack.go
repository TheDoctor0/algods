package datastructures

/*ListStack type*/
type ListStack struct {
	values []interface{}
	size   int
}

/*NewListStack creates list stack*/
func NewListStack() *ListStack {
	return &ListStack{}
}

/*Len of ListStack (current)*/
func (stack *ListStack) Len() int {
	return stack.size
}

/*Push new element to ListStack*/
func (stack *ListStack) Push(value interface{}) {
	stack.values = append(stack.values, value)
	stack.size++
}

/*Pop last element from ListStack*/
func (stack *ListStack) Pop() (value interface{}) {
	if stack.Len() == 0 {
		return nil
	}

	stack.size--
	value = stack.values[stack.size]
	stack.values[stack.size] = nil

	return value
}
