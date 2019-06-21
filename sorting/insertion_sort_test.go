package sorting_test

import (
	"testing"

	"github.com/TheDoctor0/algods/sorting"
	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	data := []int{6, 5, 3, 1, 8, 7, 2, 4}
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8}

	sorting.InsertionSort(data)

	assert.Equal(t, data, sorted, "Array should be sorted.")
}
