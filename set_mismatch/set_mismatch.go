package set_mismatch

func findErrorNums(nums []int) []int {
	m := make(map[int]bool, len(nums))
	v := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		v[1] ^= i + 1
		if m[nums[i]] {
			v[0] = nums[i]
		} else {
			m[nums[i]] = true
			v[1] ^= nums[i]
		}
	}

	return v
}

func findErrorNums2(nums []int) []int {
	m := make(map[int]bool, len(nums))
	v := make([]int, 2)

	for _, num := range nums {
		if m[num] {
			v[0] = num
		}

		m[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			v[1] = i
			break
		}
	}

	return v
}
