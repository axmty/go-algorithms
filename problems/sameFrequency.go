package problems

import "strconv"

// SameFrequency determines if integers n and m have the same frequency of digits.
// It uses the function in the same package.
func SameFrequency(n, m int) bool {
	return ValidAnagram(strconv.Itoa(n), strconv.Itoa(m))
}
