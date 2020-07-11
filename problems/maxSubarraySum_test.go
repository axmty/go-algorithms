package problems

import "testing"

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
