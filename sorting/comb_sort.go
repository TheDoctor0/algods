package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*CombSort implementation*/
func CombSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	gap := length

	for gap > 1 {
		gap = (gap * 10) / 13

		for i := 0; i+gap < length; i++ {
			if array[i] > array[i+gap] {
				utils.Swap(array, i, i+gap)
			}
		}
	}

	return array
}
