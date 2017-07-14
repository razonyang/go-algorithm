package two_sum

import "testing"

type testCase struct {
	nums   []int
	target int
	v      []int
}

var (
	testCases = []testCase{
		testCase{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			v:      []int{0, 1},
		},
		testCase{
			nums:   []int{1, 2, 3, 4},
			target: 6,
			v:      []int{1, 3},
		},
		testCase{
			nums:   []int{2, 7, 11, 15},
			target: 18,
			v:      []int{1, 2},
		},
	}
)

func TestTwoSum(t *testing.T) {
	for _, test := range testCases {
		if v := twoSum(test.nums, test.target); !isEqual(v, test.v) {
			t.Errorf("expect %v, got %v\n", test.v, v)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range testCases {
		if v := twoSum2(test.nums, test.target); !isEqual(v, test.v) {
			t.Errorf("expect %v, got %v\n", test.v, v)
		}
	}
}

func TestTwoSum3(t *testing.T) {
	for _, test := range testCases {
		if v := twoSum3(test.nums, test.target); !isEqual(v, test.v) {
			t.Errorf("expect %v, got %v\n", test.v, v)
		}
	}
}

func isEqual(v1, v2 []int) bool {
	if len(v1) != 2 || len(v1) != len(v2) {
		return false
	}

	return (v1[0] == v2[0] && v1[1] == v2[1]) || (v1[0] == v2[1] && v1[1] == v2[0])
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			_ = twoSum(test.nums, test.target)
		}
	}
}

func BenchmarkTwoSumParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range testCases {
				_ = twoSum(test.nums, test.target)
			}
		}
	})
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			_ = twoSum2(test.nums, test.target)
		}
	}
}

func BenchmarkTwoSum2Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range testCases {
				_ = twoSum2(test.nums, test.target)
			}
		}
	})
}

func BenchmarkTwoSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			_ = twoSum3(test.nums, test.target)
		}
	}
}

func BenchmarkTwoSum3Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, test := range testCases {
				_ = twoSum3(test.nums, test.target)
			}
		}
	})
}
