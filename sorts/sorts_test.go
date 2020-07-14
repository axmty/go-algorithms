package sorts

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestBubble(t *testing.T) {
	runTest(t, Bubble)
}

func TestInsertion(t *testing.T) {
	runTest(t, Insertion)
}

func TestMerge(t *testing.T) {
	runTest(t, Merge)
}

func TestSelection(t *testing.T) {
	runTest(t, Selection)
}

func runTest(t *testing.T, sortFunc func([]int)) {
	funcName := getFuncName(sortFunc)
	cases := []struct {
		arr  []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{2}, []int{2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
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

func getFuncName(sortFunc func([]int)) string {
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(sortFunc).Pointer()).Name()
	return fullFuncName[strings.LastIndex(fullFuncName, ".")+1:]
}
