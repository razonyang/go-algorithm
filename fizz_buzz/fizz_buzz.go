package fizz_buzz

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ret[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			ret[i-1] = "Buzz"
		} else if i%3 == 0 {
			ret[i-1] = "Fizz"
		} else {
			ret[i-1] = strconv.Itoa(i)
		}
	}

	return ret
}

func fizzBuzz2(n int) []string {
	ret := make([]string, n)
	three := 3
	five := 5
	fifteen := 15

	for i := 1; i <= n; i++ {
		if i == fifteen {
			ret[i-1] = "FizzBuzz"
			fifteen += 15
			five += 5
			three += 3
		} else if i == five {
			ret[i-1] = "Buzz"
			five += 5
		} else if i == three {
			ret[i-1] = "Fizz"
			three += 3
		} else {
			ret[i-1] = strconv.Itoa(i)
		}
	}

	return ret
}
