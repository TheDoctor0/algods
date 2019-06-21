package utils_test

import (
	"github.com/TheDoctor0/algods/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	data := []int{1, 2}
	swapped := []int{2, 1}

	utils.Swap(data, 0, 1)

	assert.Equal(t, data, swapped)
}
