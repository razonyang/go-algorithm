package set_mismatch

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums []int
	v    []int
}

var (
	testCases = []testCase{
		testCase{
			nums: []int{1, 1},
			v:    []int{1, 2},
		},
		testCase{
			nums: []int{2, 2},
			v:    []int{2, 1},
		},
		testCase{
			nums: []int{3, 2, 2},
			v:    []int{2, 1},
		},
		testCase{
			nums: []int{1, 2, 2, 4},
			v:    []int{2, 3},
		},
	}
)

func TestFindErrorNums(t *testing.T) {
	for _, test := range testCases {
		if v := findErrorNums(test.nums); !reflect.DeepEqual(v, test.v) {
			t.Errorf("nums: %v, expected %v, got %v\n", test.nums, v, test.v)
		}
	}
}

func TestFindErrorNums2(t *testing.T) {
	for _, test := range testCases {
		if v := findErrorNums2(test.nums); !reflect.DeepEqual(v, test.v) {
			t.Errorf("nums: %v, expected %v, got %v\n", test.nums, v, test.v)
		}
	}
}
