package problems

import "testing"

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
