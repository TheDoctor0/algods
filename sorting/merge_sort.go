package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*MergeSort implementation*/
func MergeSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	return mergeSort(array)
}

func mergeSort(array []int) []int {
	length := len(array)

	if length < 2 {
		return array
	}

	half := length / 2

	return merge(mergeSort(array[:half]), mergeSort(array[half:]))
}

func merge(left []int, right []int) []int {
	length, i, j := len(left)+len(right), 0, 0
	merged := make([]int, length, length)

	for k := 0; k < length; k++ {
		if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}

	return merged
}
