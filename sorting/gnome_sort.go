package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*GnomeSort implementation*/
func GnomeSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	index := 0

	for index < length {
		if index == 0 || array[index] >= array[index-1] {
			index++
		} else {
			utils.Swap(array, index, index-1)

			index--
		}
	}

	return array
}
