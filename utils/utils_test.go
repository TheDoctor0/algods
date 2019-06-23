package utils_test

import (
	"reflect"
	"testing"

	"github.com/TheDoctor0/algods/utils"
	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	input := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}
	swapped := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	utils.Swap(input, 3, 4)

	assert.Equal(t, input, swapped, "Values for given indexes should be swapped.")
}

func TestCopyAndLength(t *testing.T) {
	input := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}

	cloned, length := utils.CopyAndGetLength(input)

	assert.True(t, reflect.DeepEqual(input, cloned), "Cloned array should be identical.")
	assert.Equal(t, len(input), length, "Cloned array should have same length.")
}

func TestShuffle(t *testing.T) {
	input := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}
	same := []int{1, 2, 3, 5, 4, 6, 7, 8, 9}

	utils.Shuffle(input)

	assert.NotEqual(t, same, input, "Shuffled array should not be the same as before shuffle.")
}
