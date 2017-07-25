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
	}
)

func TestReplaceWords(t *testing.T) {
	for _, test := range testCases {
		if v := replaceWords(test.dict, test.sentence); v != test.v {
			t.Errorf("expect %q, got %q\n", test.v, v)
		}
	}
}
