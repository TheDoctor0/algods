package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

// Bubble sort - Bubbles up bigger values to the right side.
// Runtime: O(n^2)
//
// @param []int
func BubbleSort(array []int) {
	if len(array) <= 1 {
		return
	}

	swapped := true

	for i := len(array); i > 0; i-- {
		swapped = false

		for j := 0; j+1 < i; j++ {
			if array[j] > array[j+1] {
				swapped = utils.Swap(array, j+1, j)
			}
		}

		if !swapped {
			break
		}
	}
}
