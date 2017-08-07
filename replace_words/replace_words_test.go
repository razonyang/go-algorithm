package replace_words

import "testing"

type testCase struct {
	dict     []string
	sentence string
	v        string
}

var (
	dict1 = []string{"cat", "bat", "rat"}
	dict2 = []string{"cat", "bat", "rat", "ra"}

	testCases = []testCase{
		testCase{
			dict:     dict1,
			sentence: "the cattle was rattled by the battery",
			v:        "the cat was rat by the bat",
		},
		testCase{
			dict:     dict2,
			sentence: "the cattle was rattled by the battery",
			v:        "the cat was ra by the bat",
		},
		testCase{
			dict:     []string{"a", "aa", "aaa", "aaaa"},
			sentence: "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa",
			v:        "a a a a a a a a bbb baba a",
		},
	}
)

func TestReplaceWords(t *testing.T) {
	for _, test := range testCases {
		if v := replaceWords(test.dict, test.sentence); v != test.v {
			t.Errorf("expected %q, got %q\n", test.v, v)
		}
	}
}
