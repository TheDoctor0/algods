package utils_test

import (
	"testing"

	"github.com/TheDoctor0/algods/utils"
	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	data := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}
	swapped := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	utils.Swap(data, 3, 4)

	assert.Equal(t, data, swapped, "Values for given indexes should be swapped.")
}
