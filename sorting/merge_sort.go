package sorting

/*
 * MergeSort - uses divide and conquer paradigm.
 *
 * Runtime: O(n log n)
 */
func MergeSort(array []int) {
    if len(array) < 2 {
        return
    }

    copy(array, SplitSort(array))
}

/*
 * SplitSort - recursively splits array in half, sorts these split elements and merges them back.
 */
func SplitSort(array []int) []int {
    size := len(array)

    if size < 2 {
        return array
    }

    half := size / 2

    return Merge(SplitSort(array[:half]), SplitSort(array[half:]))
}

/*
 * Merge - merge two arrays in ascending order.
 *
 * Runtime: O(a+b)
 */
func Merge(left []int, right []int) []int {
    size, i, j := len(left)+len(right), 0, 0
    merged := make([]int, size, size)

    for k := 0; k < size; k++ {
        if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
            merged[k] = left[i]
            i++
        } else {
            merged[k] = right[j]
            j++
        }
    }

    return merged
}
