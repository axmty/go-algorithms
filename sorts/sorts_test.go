package sorts

import (
	"reflect"
	"testing"
)

var sortTestCases = []struct {
	arr  []int
	want []int
}{
	{[]int{}, []int{}},
	{[]int{2}, []int{2}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{6, 5, 3, 1, 5, 8, 7, 7, 2, 4}, []int{1, 2, 3, 4, 5, 5, 6, 7, 7, 8}},
}

func TestBubble(t *testing.T) {
	for _, c := range sortTestCases {
		arg := append([]int(nil), c.arr...)
		Bubble(c.arr)
		if !reflect.DeepEqual(c.arr, c.want) {
			t.Errorf("Bubble(%v) => %v, want %v", arg, c.arr, c.want)
		}
	}
}
