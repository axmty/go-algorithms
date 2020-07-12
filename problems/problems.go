package problems

import (
	"strconv"
)

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
	start, end := 0, len(sorted)-1
	for start < end {
		avg := (float32(sorted[start]) + float32(sorted[end])) / 2
		if avg == target {
			break
		} else if avg < target {
			start++
		} else {
			end--
		}
	}
	return start < end
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

// MaxSubArraySum finds the maximum sum of a subarray in array arr, with length n.
func MaxSubArraySum(arr []int, n int) (int, bool) {
	if n <= 0 || len(arr) < n {
		return 0, false
	}
	max := 0
	end := 0
	for end = range arr[:n] {
		max += arr[end]
	}
	sum := max
	for start, end := 1, end+1; end < len(arr); start, end = start+1, end+1 {
		sum -= arr[start]
		sum += arr[end]
		if sum > max {
			max = sum
		}
	}
	return max, true
}

// MinSubArrayLen returns the minimal length of a contiguous subarray of which
// the sum is greater than or equal to n.
func MinSubArrayLen(arr []int, n int) int {
	start, end, min, sum := 0, 0, 0, 0
	for {
		if sum < n && end < len(arr) {
			sum += arr[end]
			end++
		} else if sum >= n {
			if min == 0 || end-start < min {
				min = end - start
			}
			sum -= arr[start]
			start++

		} else {
			break
		}
	}
	return min
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
