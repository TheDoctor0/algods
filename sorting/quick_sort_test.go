package sorting_test

import (
	"testing"

	"github.com/TheDoctor0/algods/sorting"
	"github.com/stretchr/testify/assert"
)

func TestQuickSortForSingleValue(t *testing.T) {
	single := []int{0}

	assert.Equal(t, single, sorting.QuickSort(single), "Array sorting should be skipped for less than 2 elements.")
}

func TestQuickSort(t *testing.T) {
	data := []int{6, 9, 5, 3, 1, 8, 7, 2, 4}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, sorted, sorting.QuickSort(data), "Array should be sorted.")
}
