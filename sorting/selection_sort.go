package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/* SelectionSort implementation */
func SelectionSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
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

	return array
}
