package problems

// CountUniqueValues counts the unique values in the array arr.
func CountUniqueValues(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	n := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != n {
			count++
			n = arr[i]
		}
	}
	return count
}
