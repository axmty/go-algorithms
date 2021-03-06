package problems

import (
	"strconv"
	"strings"
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

// CapitalizeFirst returns the strings in arr with their first letter capitalized.
func CapitalizeFirst(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}
	return append([]string{strings.Title(arr[0])}, CapitalizeFirst(arr[1:])...)
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

// Factorial returns !n.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fib returns nth number in the Fibonacci sequence.
func Fib(n int) int {
	var fn func(int, int) int
	fn = func(curr, next int) int {
		if n == 1 {
			return curr
		}
		n--
		return fn(next, curr+next)
	}
	return fn(1, 1)
}

// FindLongestSubstring returns the length of the longest substring with all distinct characters.
func FindLongestSubstring(s string) int {
	lookup := make(map[byte]bool)
	longest := 0
	start, end := 0, 0
	for {
		if end < len(s) && !lookup[s[end]] {
			lookup[s[end]] = true
			end++
			if end-start > longest {
				longest = end - start
			}
		} else if end < len(s) && lookup[s[end]] {
			lookup[s[start]] = false
			start++
		} else {
			break
		}
	}
	return longest
}

// Flatten recursively flattens array or slice elements in arr.
func Flatten(arr []interface{}) []interface{} {
	if len(arr) == 0 {
		return []interface{}{}
	}
	switch elem := arr[0].(type) {
	case []interface{}:
		return append(Flatten(elem), Flatten(arr[1:])...)
	default:
		return append([]interface{}{elem}, Flatten(arr[1:])...)
	}
}

// IsPalindrome if string s is a palindrome.
func IsPalindrome(s string) bool {
	var fn func([]rune) bool
	fn = func(runes []rune) bool {
		if len(runes) <= 1 {
			return true
		}
		return runes[0] == runes[len(runes)-1] && fn(runes[1:len(runes)-1])
	}
	return fn([]rune(s))
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

// NestedEvenSum returns the sum of all nested even integers in arr.
func NestedEvenSum(arr []interface{}) int {
	if len(arr) == 0 {
		return 0
	}
	v := 0
	switch elem := arr[0].(type) {
	case []interface{}:
		v = NestedEvenSum(elem)
	case int:
		v = elem
	}
	return v + NestedEvenSum(arr[1:])
}

// Pow returns n power pow.
func Pow(n, pow int) int {
	if pow == 0 {
		return 1
	}
	return n * Pow(n, pow-1)
}

// Product returns the product of all the elements of arr.
func Product(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var fn func(int) int
	fn = func(i int) int {
		if i == len(arr) {
			return 1
		}
		return arr[i] * fn(i+1)
	}
	return fn(0)
}

// Reverse returns s in reverse.
func Reverse(s string) string {
	var fn func([]rune) string
	fn = func(runes []rune) string {
		switch len(runes) {
		case 0:
			return ""
		case 1:
			return string(runes[0])
		default:
			return Reverse(string(runes[1:])) + string(runes[0])
		}
	}
	return fn([]rune(s))
}

// SameFrequency determines if integers n and m have the same frequency of digits.
// It uses the function in the same package.
func SameFrequency(n, m int) bool {
	return ValidAnagram(strconv.Itoa(n), strconv.Itoa(m))
}

// SomeRecursive determines if at least one value in arr verifies the given predicate pre.
func SomeRecursive(arr []interface{}, predicate func(interface{}) bool) bool {
	if len(arr) == 0 {
		return false
	}
	return predicate(arr[0]) || SomeRecursive(arr[1:], predicate)
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
