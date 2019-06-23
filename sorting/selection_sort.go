package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*SelectionSort implementation*/
func SelectionSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	for i := 0; i < length; i++ {
		selection := i

		for j := i + 1; j < length; j++ {
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
