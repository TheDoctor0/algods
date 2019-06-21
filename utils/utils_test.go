package utils_test

import (
	"testing"

	"github.com/TheDoctor0/algods/utils"
	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	data := []int{1, 2}
	swapped := []int{2, 1}

	utils.Swap(data, 0, 1)

	assert.Equal(t, data, swapped, "Values for given indexes should be swapped.")
	assert.False(t, utils.Swap(data, 1, 1), "Same indexes cannot be swapped.")
	assert.False(t, utils.Swap(data, 0, len(data)), "Indexes cannot be greater or equal array length.")
	assert.False(t, utils.Swap(data, 0, -1), "Indexes cannot be smaller than 0.")
}
