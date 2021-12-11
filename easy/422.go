package easy

func FindAllDuplicate(arr []int) []int {
	m := map[int]int{}
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] == 0 {
			m[arr[i]] = 1
		} else {
			m[arr[i]] += 1
		}
	}

	res := []int{}
	for key, val := range m {
		if val == 2 {
			res = append(res, key)
		}
	}

	return res
}
