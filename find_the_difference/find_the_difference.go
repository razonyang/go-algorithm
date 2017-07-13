package find_the_difference

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]] += 1
	}

	var b byte
	for i := 0; i < len(t); i++ {
		if count, ok := m[t[i]]; !ok || count <= 0 {
			b = t[i]
			break
		} else {
			m[t[i]] -= 1
		}
	}

	return b
}

func findTheDifference2(s string, t string) byte {
	var b byte
	for i := 0; i < len(s); i++ {
		b ^= s[i] ^ t[i]
	}
	b ^= t[len(t)-1]

	return b
}
