package sorts

// Bubble sorts arr using bubble sort algorithm.
func Bubble(arr []int) {
	var swapped bool = true
	for n := len(arr) - 1; swapped; n-- {
		swapped = false
		for j := 0; j < n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}
}

// Selection sorts arr using selection sort algorithm.
func Selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
