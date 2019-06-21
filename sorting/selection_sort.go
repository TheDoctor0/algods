package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

// Selection sort - Look for smaller numbers on the right side.
// Runtime: O(n^2)
//
// @param []int
func SelectionSort(array []int) {
	if len(array) <= 1 {
		return
	}

	for i := 0; i < len(array); i++ {
		selection := i

		for j := i + 1; j < len(array); j++ {
			if array[selection] > array[j] {
				selection = j
			}
		}

		if selection != i {
			utils.Swap(array, selection, i)
		}
	}
}
