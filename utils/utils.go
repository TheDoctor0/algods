package utils

func Swap(array []int, from int, to int) {
	array[from], array[to] = array[to], array[from]
}
