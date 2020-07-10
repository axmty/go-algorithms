package problems

import "testing"

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
