package longest_common_prefix

import "testing"

type testCase struct {
	strs   []string
	prefix string
}

var (
	testCases = []testCase{
		testCase{[]string{}, ""},
		testCase{[]string{"foo"}, "foo"},
		testCase{[]string{"foo", "bar"}, ""},
		testCase{[]string{"foo", "football"}, "foo"},
		testCase{[]string{"foo", "flush"}, "f"},
		testCase{[]string{"foo", "football", "forward"}, "fo"},
		testCase{[]string{"foo", "football", "forward", "bar"}, ""},
	}
)

func TestLongestCommonPrefix(t *testing.T) {
	for _, v := range testCases {
		if prefix := longestCommonPrefix(v.strs); v.prefix != prefix {
			t.Errorf("%v longest common prefix %q, got %q", v.strs, v.prefix, prefix)
		}
	}
}
