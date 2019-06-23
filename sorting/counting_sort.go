package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*CountingSort implementation*/
func CountingSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	k := getK(array)
	counts := make([]int, k)

	for i := 0; i < length; i++ {
		counts[array[i]]++
	}

	index := 0

	for i, c := range counts {
		for ; c > 0; c-- {
			array[index] = i
			index++
		}
	}

	return array
}

func getK(array []int) int {
	max := array[0]

	for _, value := range array {
		if value > max {
			max = value
		}
	}

	return max + 1
}
