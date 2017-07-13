package find_the_difference

import "testing"

type testCase struct {
	s string
	t string
	v byte
}

var (
	testCases = []testCase{
		testCase{s: "a", t: "aa", v: 'a'},
		testCase{s: "ab", t: "abb", v: 'b'},
		testCase{s: "ab", t: "cab", v: 'c'},
		testCase{s: "aab", t: "aaba", v: 'a'},
	}
)

func TestFindTheDifference(t *testing.T) {
	for _, test := range testCases {
		if v := findTheDifference(test.s, test.t); v != test.v {
			t.Errorf("%q and %q difference %q, got %q\n", test.s, test.t, test.v, v)
		}
	}
}

func TestFindTheDifference2(t *testing.T) {
	for _, test := range testCases {
		if v := findTheDifference2(test.s, test.t); v != test.v {
			t.Errorf("%q and %q difference %q, got %q\n", test.s, test.t, test.v, v)
		}
	}
}
