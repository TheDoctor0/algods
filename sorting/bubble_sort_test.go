package sorting_test

import (
	"testing"

	"github.com/TheDoctor0/algods/sorting"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSortForSingleValue(t *testing.T) {
	single := []int{0}

	assert.Equal(t, single, sorting.BubbleSort(single), "Array sorting should be skipped for less than 2 elements.")
}

func TestBubbleSort(t *testing.T) {
	input := []int{6, 9, 5, 3, 1, 8, 7, 2, 4}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, sorted, sorting.BubbleSort(input), "Array should be sorted.")
}
