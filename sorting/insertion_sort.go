package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*InsertionSort implementation*/
func InsertionSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				utils.Swap(array, j, i)
			}
		}
	}

	return array
}
