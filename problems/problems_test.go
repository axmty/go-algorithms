package problems

import (
	"reflect"
	"testing"
)

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

func TestCapitalizeFirst(t *testing.T) {
	cases := []struct {
		arr  []string
		want []string
	}{
		{[]string{"car", "taco", "banana"}, []string{"Car", "Taco", "Banana"}},
	}
	for _, c := range cases {
		got := CapitalizeFirst(c.arr)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("CapitalizeFirst(%v) == %v, want %v", c.arr, got, c.want)
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

func TestFactorial(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{4, 24},
		{7, 5040},
	}
	for _, c := range cases {
		got := Factorial(c.n)
		if got != c.want {
			t.Errorf("Factorial(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}

func TestFib(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{4, 3},
		{10, 55},
		{28, 317811},
	}
	for _, c := range cases {
		got := Fib(c.n)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.n, got, c.want)
		}
	}
}

func TestFindLongestSubstring(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"rithmschool", 7},
		{"thisisawesome", 6},
		{"thecatinthehat", 7},
		{"bbbbbb", 1},
		{"longestsubstring", 8},
		{"thisishowwedoit", 6},
	}
	for _, c := range cases {
		got := FindLongestSubstring(c.s)
		if got != c.want {
			t.Errorf("FindLongestSubstring(%q) == %d, want %d", c.s, got, c.want)
		}
	}
}

func TestFlatten(t *testing.T) {
	cases := []struct {
		arr  []interface{}
		want []interface{}
	}{
		// [1 2 3 [4 5]] => [1 2 3 4 5]
		{
			[]interface{}{1, 2, 3, []interface{}{4, 5}},
			[]interface{}{1, 2, 3, 4, 5},
		},
		// [1 [2 [3 4] [[5]]]] => [1 2 3 4 5]
		{
			[]interface{}{1, []interface{}{2, []interface{}{3, 4}, []interface{}{[]interface{}{5}}}},
			[]interface{}{1, 2, 3, 4, 5},
		},
	}
	for _, c := range cases {
		got := Flatten(c.arr)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Flatten(%v) == %v, want %v", c.arr, got, c.want)
			continue
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s    string
		want bool
	}{
		{"awesome", false},
		{"foobar", false},
		{"tacocat", true},
		{"taccat", true},
		{"amanaplanacanalpanama", true},
		{"amanaplanacanalpandemonium", false},
		{"维基百百基维", true},
		{"维基百维基百", false},
	}
	for _, c := range cases {
		got := IsPalindrome(c.s)
		if got != c.want {
			t.Errorf("IsPalindrome(%q) == %t, want %t", c.s, got, c.want)
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

func TestMaxSubArraySum(t *testing.T) {
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
		got, ok := MaxSubArraySum(c.arr, c.n)
		if got != c.out && ok != c.ok {
			t.Errorf("MaxSubarraySum(%v, %d) == (%d, %t), want (%d, %t)",
				c.arr, c.n, got, ok, c.out, c.ok)
		}
	}
}

func TestMinSubArrayLen(t *testing.T) {
	cases := []struct {
		arr  []int
		n    int
		want int
	}{
		{[]int{}, 5, 0},
		{[]int{3}, 4, 0},
		{[]int{3}, 3, 1},
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{2, 1, 6, 5, 4}, 9, 2},
		{[]int{3, 1, 7, 11, 2, 9, 8, 21, 62, 33, 19}, 52, 1},
		{[]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 39, 3},
		{[]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 55, 5},
		{[]int{4, 3, 3, 8, 1, 2, 3}, 11, 2},
		{[]int{1, 4, 16, 22, 5, 7, 8, 9, 10}, 95, 0},
	}
	for _, c := range cases {
		got := MinSubArrayLen(c.arr, c.n)
		if got != c.want {
			t.Errorf("MinSubArrayLen(%v, %d) == %d, want %d", c.arr, c.n, got, c.want)
		}
	}
}

func TestNestedEvenSum(t *testing.T) {
	cases := []struct {
		arr  []interface{}
		want int
	}{
		// [1 'a' 3 ['b' 5]] => 9
		{[]interface{}{1, 'a', 3, []interface{}{'b', 5}}, 9},
		// [[1 2]] => 3
		{[]interface{}{[]interface{}{1, 2}}, 3},
	}
	for _, c := range cases {
		got := NestedEvenSum(c.arr)
		if got != c.want {
			t.Errorf("NestedEvenSum(%v) == %d, want %d", c.arr, got, c.want)
			continue
		}
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		n    int
		pow  int
		want int
	}{
		{2, 0, 1},
		{2, 2, 4},
		{2, 4, 16},
	}
	for _, c := range cases {
		got := Pow(c.n, c.pow)
		if got != c.want {
			t.Errorf("Pow(%d, %d) == %d, want %d", c.n, c.pow, got, c.want)
		}
	}
}

func TestProduct(t *testing.T) {
	cases := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 10}, 60},
	}
	for _, c := range cases {
		got := Product(c.arr)
		if got != c.want {
			t.Errorf("Product(%v) == %d, want %d", c.arr, got, c.want)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		s    string
		want string
	}{
		{"awesome", "emosewa"},
		{"rithmschool", "loohcsmhtir"},
		{"维基百科:关于中文维基百科", "科百基维文中于关:科百基维"},
		{"维", "维"},
	}
	for _, c := range cases {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
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

func TestSomeRecursive(t *testing.T) {
	isOdd := func(i interface{}) bool {
		return i.(int)%2 == 1
	}
	moreThanTen := func(i interface{}) bool {
		return i.(int) > 10
	}
	cases := []struct {
		arr      []interface{}
		funcName string
		pre      func(interface{}) bool
		want     bool
	}{
		{[]interface{}{1, 2, 3, 4}, "isOdd", isOdd, true},
		{[]interface{}{4, 6, 8, 9}, "isOdd", isOdd, true},
		{[]interface{}{4, 6, 8}, "isOdd", isOdd, false},
		{[]interface{}{4, 6, 8}, "isOdd", moreThanTen, false},
	}
	for _, c := range cases {
		got := SomeRecursive(c.arr, c.pre)
		if got != c.want {
			t.Errorf("SomeRecursive(%v, %s) == %t, want %t",
				c.arr, c.funcName, got, c.want)
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
