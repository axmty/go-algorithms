package problems

// MaxSubarraySum finds the maximum sum of a subarray in array arr, with length n.
func MaxSubarraySum(arr []int, n int) (int, bool) {
	if n <= 0 || len(arr) < n {
		return 0, false
	}
	max := 0
	end := 0
	for end = range arr[:n] {
		max += arr[end]
	}
	sum := max
	for begin, end := 1, end+1; end < len(arr); begin, end = begin+1, end+1 {
		sum -= arr[begin]
		sum += arr[end]
		if sum > max {
			max = sum
		}
	}
	return max, true
}
