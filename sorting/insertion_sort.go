package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

// Insertion sort - Look for bigger numbers on the left side.
// Runtime: O(n^2)
//
// @param []int
// @return bool
func InsertionSort(array []int) bool {
	if len(array) <= 1 {
		return false
	}

	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				utils.Swap(array, j, i)
			}
		}
	}

	return true
}
