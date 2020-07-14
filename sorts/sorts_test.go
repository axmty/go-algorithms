package sorts

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	cases := []struct {
		arr  []int
		want []int
	}{
		{[]int{6, 5, 3, 1, 5, 8, 7, 7, 2, 4}, []int{1, 2, 3, 4, 5, 5, 6, 7, 7, 8}},
	}
	for _, c := range cases {
		arg := append([]int(nil), c.arr...)
		Bubble(c.arr)
		if !reflect.DeepEqual(c.arr, c.want) {
			t.Errorf("Bubble(%v) => %v, want %v", arg, c.arr, c.want)
		}
	}
}
