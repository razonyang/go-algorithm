package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = true
	}

	var v []int
	for _, num := range nums2 {
		if m[num] {
			v = append(v, num)
			delete(m, num)
		}
	}

	return v
}
