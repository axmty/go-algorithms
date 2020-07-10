package problems

import "testing"

func TestSameFrequencyUsingValidAnagram(t *testing.T) {
	testSameFrequency(t, SameFrequencyUsingValidAnagram)
}

func testSameFrequency(t *testing.T, fn func(int, int) bool) {
	cases := []struct {
		n, m int
		want bool
	}{
		{182, 281, true},
		{34, 14, false},
		{22, 222, false},
	}
	for _, c := range cases {
		got := fn(c.n, c.m)
		if got != c.want {
			t.Errorf("SameFrequency(%d, %d) == %t, want %t", c.n, c.m, got, c.want)
		}
	}
}
