package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

// MergeSort - uses divide and conquer paradigm.
// Runtime: O(n log n)
//
// @param []int
func MergeSort(array []int) {
	if len(array) < 2 {
		return
	}

	copy(array, SplitSort(array))
}

// SplitSort - recursively splits array in half, sorts these split elements and merges them back.
//
// @param []int
// @return []int
func SplitSort(array []int) []int {
	size := len(array)

	if size < 2 {
		return array
	} else if size == 2 {
		if array[0] > array[1] {
			utils.Swap(array, 0, 1)
		}

		return array
	}

	half := size / 2

	return Merge(SplitSort(array[:half]), SplitSort(array[half:]))
}

// Merge - merge two arrays in ascending order.
// Runtime: O(a+b)
//
// @param []int
// @param []int
// @return []int
func Merge(left []int, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		}

		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged
}
