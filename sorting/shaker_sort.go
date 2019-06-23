package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*ShakerSort implementation*/
func ShakerSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	start := 0
	end := length - 1

	for start <= end {
		newStart := end
		newEnd := start
		swap := false

		for i := start; i < end; i++ {
			if array[i] > array[i+1] {
				utils.Swap(array, i, i+1)

				swap = true
				newEnd = i
			}
		}

		if !swap {
			break
		}

		end = newEnd
		swap = false

		for i := end; i >= start; i-- {
			if array[i] > array[i+1] {
				utils.Swap(array, i, i+1)

				swap = true
				newStart = i
			}
		}

		if !swap {
			break
		}

		start = newStart
	}

	return array
}
