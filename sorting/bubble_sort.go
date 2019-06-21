package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

func BubbleSort(array []int) {
	swapped := true

	for i := len(array); i > 0; i-- {
		swapped = false

		for j := 0; j+1 < i; j++ {
			if array[j] > array[j+1] {
				utils.Swap(array, j+1, j)

				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
