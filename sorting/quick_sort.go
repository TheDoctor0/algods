package sorting

import (
    "math/rand"
    "time"

    "github.com/TheDoctor0/algods/utils"
)

/*
 * QuickSort - quick sort with array shuffle to avoid edgy cases.
 *
 * Runtime: O(n log n)
 */
func QuickSort(array []int) {
    if len(array) < 2 {
        return
    }

    Shuffle(array)

    copy(array, QuickSortPartition(array, 0, len(array)-1))
}

/*
 * Shuffle - shuffle elements in array (Fisher–Yates algorithm).
 *
 * Runtime: O(n)
 */
func Shuffle(array []int) {
    rand.Seed(time.Now().UnixNano())

    for i := len(array) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)

        utils.Swap(array, i, j)
    }
}

/*
 * QuickSortPartition - uses recursive linear time partitioning.
 */
func QuickSortPartition(array []int, low int, high int) []int {
    if low < high {
        partitionIndex := Partition(array, low, high)

        QuickSortPartition(array, low, partitionIndex-1)
        QuickSortPartition(array, partitionIndex+1, high)
    }

    return array
}

/*
 * Partition - Chooses a pivot and re-arrange the array against it.
 *
 * Runtime: O(n)
 */
func Partition(array []int, low int, high int) int {
    pivotIndex := low

    for i := low + 1; i <= high; i++ {
        if array[i] < array[low] {
            pivotIndex++

            utils.Swap(array, i, pivotIndex)
        }
    }

    utils.Swap(array, low, pivotIndex)

    return pivotIndex
}
