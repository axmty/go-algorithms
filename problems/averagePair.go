package problems

// AveragePair determines if there is a pair of values in arr where the average of the pair equals
// the target average.
func AveragePair(arr []int, target float32) bool {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	begin, end := 0, len(sorted)-1
	for begin < end {
		avg := (float32(sorted[begin]) + float32(sorted[end])) / 2
		if avg == target {
			break
		} else if avg < target {
			begin++
		} else {
			end--
		}
	}
	return begin < end
}
