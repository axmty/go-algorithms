package sorts

import (
	"reflect"
	"runtime"
	"testing"
)

func TestBubble(t *testing.T) {
	runTest(t, Bubble)
}

func TestSelection(t *testing.T) {
	runTest(t, Selection)
}

func runTest(t *testing.T, sortFunc func([]int)) {
	funcName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
	cases := []struct {
		arr  []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{6, 5, 3, 1, 5, 8, 7, 7, 2, 4}, []int{1, 2, 3, 4, 5, 5, 6, 7, 7, 8}},
	}
	for _, c := range cases {
		arg := append([]int(nil), c.arr...)
		sortFunc(c.arr)
		if !reflect.DeepEqual(c.arr, c.want) {
			t.Errorf("%s(%v) => %v, want %v", funcName, arg, c.arr, c.want)
		}
	}
}
