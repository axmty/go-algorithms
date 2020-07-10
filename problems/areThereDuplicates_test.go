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
