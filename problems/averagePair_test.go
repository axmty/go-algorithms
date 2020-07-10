package problems

import "testing"

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
