package intersection_of_two_arrays

import (
	"sort"
	"testing"
)

type testCase struct {
	nums1 []int
	nums2 []int
	v     []int
}

var (
	testCases = []testCase{
		testCase{
			nums1: []int{},
			nums2: []int{},
			v:     []int{},
		},
		testCase{
			nums1: []int{1},
			nums2: []int{},
			v:     []int{},
		},
		testCase{
			nums1: []int{1},
			nums2: []int{1},
			v:     []int{1},
		},
		testCase{
			nums1: []int{1},
			nums2: []int{2},
			v:     []int{},
		},
		testCase{
			nums1: []int{1, 2, 2, 4},
			nums2: []int{2, 2},
			v:     []int{2},
		},
		testCase{
			nums1: []int{1, 2, 2, 4},
			nums2: []int{2, 2, 4},
			v:     []int{4, 2},
		},
	}
)

func TestInterSection(t *testing.T) {
	for _, test := range testCases {
		if v := intersection(test.nums1, test.nums2); !isEqual(test.v, v) {
			t.Errorf("nums1: %v, nums2: %v, expected %v, got %v\n", test.nums1, test.nums2, test.v, v)
		}
	}
}

func isEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}

	sort.Ints(nums1)
	sort.Ints(nums2)

	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}

	return true
}
