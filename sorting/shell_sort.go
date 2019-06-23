package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

/*ShellSort implementation*/
func ShellSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	for k := length / 2; k > 0; k /= 2 {
		for i := k; i < length; i++ {
			for j := i; j >= k && array[j-k] > array[j]; j -= k {
				utils.Swap(array, j-k, j)
			}
		}
	}

	return array
}
