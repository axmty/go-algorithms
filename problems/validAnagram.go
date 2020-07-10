package problems

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
