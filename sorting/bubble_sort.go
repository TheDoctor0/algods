package sorting

func BubbleSort(input []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := len(input); i > 0; i-- {
			for j := 0; j+1 < i; j++ {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]

					swapped = true
				}
			}
		}
	}
}
