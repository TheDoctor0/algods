package utils

import (
	"math/rand"
	"time"
)

/*CopyAndGetLength - returns copy of array and its size.*/
func CopyAndGetLength(input []int) ([]int, int) {
	length := len(input)
	array := make([]int, length)
	copy(array, input)

	return array, length
}

/*Swap - swaps values for specified indexes in given array.*/
func Swap(array []int, from int, to int) {
	array[from], array[to] = array[to], array[from]
}

/*Shuffle - shuffle elements in array using Fisher–Yates algorithm.*/
func Shuffle(array []int) {
	rand.Seed(time.Now().UnixNano())

	for i := len(array) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		Swap(array, i, j)
	}
}
