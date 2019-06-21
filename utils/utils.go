package utils

// Swap - swaps values with specified indexes in given array.
//
// @param []int
// @param int
// @param int
// @return bool
func Swap(array []int, from int, to int) bool {
	if from == to || from < 0 || from >= len(array) || to < 0 || to >= len(array) {
		return false
	}

	array[from], array[to] = array[to], array[from]

	return true
}
