package easy

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
	}

	for k, _ := range m {
		if m[k] != 0 {
			return false
		}
	}

	return true
}
