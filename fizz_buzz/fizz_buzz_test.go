package fizz_buzz

import (
	"reflect"
	"testing"
)

const (
	benchmarkN = 100000
)

type testCase struct {
	n int
	v []string
}

var (
	testCases = []testCase{
		testCase{0, []string{}},
		testCase{1, []string{"1"}},
		testCase{2, []string{"1", "2"}},
		testCase{3, []string{"1", "2", "Fizz"}},
		testCase{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		testCase{15, []string{
			"1", "2", "Fizz", "4", "Buzz",
			"Fizz", "7", "8", "Fizz", "Buzz",
			"11", "Fizz", "13", "14", "FizzBuzz",
		}},
	}
)

func TestFizzBuzz(t *testing.T) {
	for _, test := range testCases {
		if v := fizzBuzz(test.n); !reflect.DeepEqual(test.v, v) {
			t.Errorf("fizzBuzz(%d), expected %v, got %v\n", test.n, test.v, v)
		}
	}
}

func TestFizzBuzz2(t *testing.T) {
	for _, test := range testCases {
		if v := fizzBuzz2(test.n); !reflect.DeepEqual(test.v, v) {
			t.Errorf("fizzBuzz(%d), expected %v, got %v\n", test.n, test.v, v)
		}
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fizzBuzz(benchmarkN)
	}
}

func BenchmarkFizzBuzzParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fizzBuzz(benchmarkN)
		}
	})
}

func BenchmarkFizzBuzz2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fizzBuzz2(benchmarkN)
	}
}

func BenchmarkFizzBuzz2Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fizzBuzz2(benchmarkN)
		}
	})
}
