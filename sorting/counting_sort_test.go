package sorting_test

import (
	"github.com/TheDoctor0/algods/sorting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSortForSingleValue(t *testing.T) {
	single := []int{0}

	assert.Equal(t, single, sorting.CountingSort(single), "Array sorting should be skipped for less than 2 elements.")
}

func TestCountingSort(t *testing.T) {
	input := []int{6, 9, 5, 3, 1, 8, 7, 2, 4}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, sorted, sorting.CountingSort(input), "Array should be sorted.")
}
