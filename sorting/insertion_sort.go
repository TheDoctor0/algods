package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

func InsertionSort(array []int) {
	if len(array) <= 1 {
		return
	}

	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j] > array[i] {
				utils.Swap(array, j, i)
			}
		}
	}
}
