package easy

func FirstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	m := make(map[string]int, 26)
	for _, char := range s {
		m[string(char)]++
	}

	for idx, char := range s {
		if m[string(char)] == 1 {
			return idx
		}
	}

	return -1
}
