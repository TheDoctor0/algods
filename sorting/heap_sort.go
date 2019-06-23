package sorting

import (
	"github.com/TheDoctor0/algods/utils"
)

type maxHeap struct {
	array []int
	size  int
}

/*HeapSort implementation*/
func HeapSort(input []int) []int {
	array, length := utils.CopyAndGetLength(input)

	if length < 2 {
		return array
	}

	return heapSort(array)
}

func buildMaxHeap(array []int) maxHeap {
	heap := maxHeap{array: array, size: len(array)}

	for i := len(array) / 2; i >= 0; i-- {
		heap.heapify(i)
	}

	return heap
}

func (heap maxHeap) heapify(i int) {
	max, left, right := i, 2*i+1, 2*i+2

	if left < heap.size && heap.array[left] > heap.array[max] {
		max = left
	}

	if right < heap.size && heap.array[right] > heap.array[max] {
		max = right
	}

	if max != i {
		utils.Swap(heap.array, max, i)

		heap.heapify(max)
	}
}

func heapSort(array []int) []int {
	heap := buildMaxHeap(array)

	for i := len(heap.array) - 1; i >= 1; i-- {
		utils.Swap(heap.array, i, 0)

		heap.size--
		heap.heapify(0)
	}

	return heap.array
}
