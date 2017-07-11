package longest_common_prefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	max := len(strs[0])

	for i := 1; i < len(strs); i++ {
		max = min(max, len(strs[i]))
		for j := 0; j < max; j++ {
			if strs[i][j] != strs[i-1][j] {
				max = j
				continue
			}
		}
	}

	return strs[0][:max]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
