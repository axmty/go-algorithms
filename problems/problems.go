package problems

import "strconv"

// AreThereDuplicates determines if there are duplicates in the params.
func AreThereDuplicates(params ...interface{}) bool {
	m := make(map[interface{}]bool)
	for _, p := range params {
		if _, exist := m[p]; exist {
			return true
		}
		m[p] = true
	}
	return false
}

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

// IsSubsequence determines if string a is a subsequence of string b.
func IsSubsequence(a, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == len(a)
}

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

// SameFrequency determines if integers n and m have the same frequency of digits.
// It uses the function in the same package.
func SameFrequency(n, m int) bool {
	return ValidAnagram(strconv.Itoa(n), strconv.Itoa(m))
}

// ValidAnagram determines if string b is an anagram of string a.
func ValidAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	lookup := make(map[rune]int)
	for _, r := range a {
		if _, ok := lookup[r]; !ok {
			lookup[r] = 1
		} else {
			lookup[r]++
		}
	}
	for _, r := range b {
		if n, ok := lookup[r]; !ok || n == 0 {
			return false
		}
		lookup[r]--
	}
	return true
}
