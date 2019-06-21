package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*
 * BubbleSort - bubbles up bigger values to the right side.
 *
 * Runtime: O(n^2)
 */
func BubbleSort(array []int) {
	if len(array) < 2 {
		return
	}

	for i := len(array); i > 0; i-- {
		swapped := false

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
