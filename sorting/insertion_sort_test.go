package sorting_test

import (
    "testing"

    "github.com/TheDoctor0/algods/sorting"
    "github.com/stretchr/testify/assert"
)

func TestInsertionSortForSingleValue(t *testing.T) {
    single := []int{0}

    sorting.InsertionSort(single)

    assert.Equal(t, []int{0}, single, "Array sorting should be skipped for less than 2 elements.")
}

func TestInsertionSort(t *testing.T) {
    data := []int{6, 9, 5, 3, 1, 8, 7, 2, 4}
    sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    sorting.InsertionSort(data)

    assert.Equal(t, sorted, data, "Array should be sorted.")
}
