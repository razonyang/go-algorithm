package search_insert_position

func searchInsert(nums []int, target int) int {
	for idx, num := range nums {
		if num >= target {
			return idx
		}
	}

	return len(nums)
}
