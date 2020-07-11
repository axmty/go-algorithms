package problems

import "testing"

func TestAreThereDuplicates(t *testing.T) {
	cases := []struct {
		arr  []interface{}
		want bool
	}{
		{[]interface{}{1, 2, 3}, false},
		{[]interface{}{1, 2, 2}, true},
		{[]interface{}{'a', 'b', 'c', 'a'}, true},
	}
	for _, c := range cases {
		got := AreThereDuplicates(c.arr...)
		if got != c.want {
			t.Errorf("AreThereDuplicates(%v) == %t, want %t", c.arr, got, c.want)
		}
	}
}

func TestAveragePair(t *testing.T) {
	cases := []struct {
		arr    []int
		target float32
		want   bool
	}{
		{[]int{1, 2, 3}, 2.5, true},
		{[]int{1, 3, 3, 5, 6, 7, 10, 12, 19}, 8, true},
		{[]int{-1, 0, 3, 4, 5, 6}, 4.1, false},
		{[]int{}, 4, false},
	}
	for _, c := range cases {
		got := AveragePair(c.arr, c.target)
		if got != c.want {
			t.Errorf("AveragePair(%v, %f) == %t, want %t", c.arr, c.target, got, c.want)
		}
	}
}

func TestCountUniqueValues(t *testing.T) {
	cases := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 1, 1, 1, 2}, 2},
		{[]int{1, 2, 3, 4, 4, 4, 7, 12, 13}, 7},
		{[]int{}, 0},
		{[]int{-2, -1, -1, 0, 1}, 4},
	}
	for _, c := range cases {
		got := CountUniqueValues(c.arr)
		if got != c.want {
			t.Errorf("CountUniqueValues(%v) == %d, want %d", c.arr, got, c.want)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"hello", "hello world", true},
		{"sing", "sting", true},
		{"abc", "abracadabra", true},
		{"abc", "acb", false},
	}
	for _, c := range cases {
		got := IsSubsequence(c.a, c.b)
		if got != c.want {
			t.Errorf("IsSubsequence(%q, %q) == %t, want %t", c.a, c.b, got, c.want)
		}
	}
}

func TestMaxSubarraySum(t *testing.T) {
	cases := []struct {
		arr []int
		n   int
		out int
		ok  bool
	}{
		{[]int{100, 200, 300, 400}, 2, 700, true},
		{[]int{1, 4, 2, 10, 23, 3, 1, 0, 20}, 4, 39, true},
		{[]int{-3, 4, 0, -2, 6, -1}, 2, 5, true},
		{[]int{3, -2, 7, -4, 1, -1, 4, -2, 1}, 2, 5, true},
		{[]int{2, 3}, 3, 0, false},
	}
	for _, c := range cases {
		got, ok := MaxSubarraySum(c.arr, c.n)
		if got != c.out && ok != c.ok {
			t.Errorf("MaxSubarraySum(%v, %d) == (%d, %t), want (%d, %t)",
				c.arr, c.n, got, ok, c.out, c.ok)
		}
	}
}

func TestSameFrequency(t *testing.T) {
	cases := []struct {
		n, m int
		want bool
	}{
		{182, 281, true},
		{34, 14, false},
		{22, 222, false},
	}
	for _, c := range cases {
		got := SameFrequency(c.n, c.m)
		if got != c.want {
			t.Errorf("SameFrequency(%d, %d) == %t, want %t", c.n, c.m, got, c.want)
		}
	}
}

func TestValidAnagram(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"", "", true},
		{"aaz", "", false},
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"awesome", "awesom", false},
	}
	for _, c := range cases {
		got := ValidAnagram(c.a, c.b)
		if got != c.want {
			t.Errorf("ValidAnagram(%q, %q) == %t, want %t", c.a, c.b, got, c.want)
		}
	}
}
