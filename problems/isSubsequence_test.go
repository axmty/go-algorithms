package problems

import "testing"

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
