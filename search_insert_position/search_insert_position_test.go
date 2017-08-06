package search_insert_position

import "testing"

type testCase struct {
	nums   []int
	target int
	v      int
}

var (
	testCases = []testCase{
		testCase{
			nums:   []int{1},
			target: 0,
			v:      0,
		},
		testCase{
			nums:   []int{1},
			target: 1,
			v:      0,
		},
		testCase{
			nums:   []int{1},
			target: 2,
			v:      1,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			v:      2,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			v:      1,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			v:      4,
		},
		testCase{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			v:      0,
		},
	}
)

func TestSearchInsert(t *testing.T) {
	for _, test := range testCases {
		if v := searchInsert(test.nums, test.target); v != test.v {
			t.Errorf("nums: %v, target: %d, expect %d, got %d\n", test.nums, test.target, test.v, v)
		}
	}
}
