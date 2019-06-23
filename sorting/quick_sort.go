package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/* QuickSort implementation */
func QuickSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	utils.Shuffle(array)

	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, low int, high int) []int {
	if low < high {
		partitionIndex := partition(array, low, high)

		quickSort(array, low, partitionIndex-1)
		quickSort(array, partitionIndex+1, high)
	}

	return array
}

func partition(array []int, low int, high int) int {
	pivotIndex := low

	for i := low + 1; i <= high; i++ {
		if array[i] < array[low] {
			pivotIndex++

			utils.Swap(array, i, pivotIndex)
		}
	}

	utils.Swap(array, low, pivotIndex)

	return pivotIndex
}
