package strings

import (
	"testing"
)

func TestStringsContainSameCharsAndLen(t *testing.T) {
	cases := []struct {
		a, b   string
		isTrue bool
	}{
		{"abc", "cab", true},
		{"dhaf", "dhra", false},
		{"abc", "abcc", false},
		{"ghjk", "jkhg", true},
	}
	for _, c := range cases {
		res := StringsContainSameCharsAndLen(c.a, c.b)
		if res != c.isTrue {
			t.Fail()
		}
	}
}
