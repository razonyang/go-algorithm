package two_sum

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		minus := target - nums[i]
		if j, ok := m[minus]; ok && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}

func twoSum3(nums []int, target int) []int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		minus := target - nums[i]
		if j, ok := m[minus]; ok {
			return []int{i, j}
		}

		m[nums[i]] = i
	}

	return []int{}
}
