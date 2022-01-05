package maps

import "testing"

func TestStringInMap(t *testing.T) {

	cases := []struct {
		m   map[int]string
		s   string
		res bool
	}{
		{
			map[int]string{
				1: "abc",
				2: "bcd",
				3: "def",
			},
			"def",
			true,
		},
		{
			map[int]string{
				100:      "dnjmnvc87",
				1000:     "dasas",
				13213123: "dgsfkldsjfkdsfj",
			},
			"dasdlasdaksld",
			false,
		},
	}

	for _, c := range cases {
		res := StringInMap(c.m, c.s)
		if res != c.res {
			t.Fail()
		}
	}
}
