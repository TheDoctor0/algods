package utils

/*
 * Swap - swaps values with specified indexes in given array.
 */
func Swap(array []int, from int, to int) {
    array[from], array[to] = array[to], array[from]
}
