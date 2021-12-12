package easy

func ContainsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			m[nums[i]] = 1
		} else {
			m[nums[i]] += 1
		}
	}

	for _, v := range m {
		if v >= 2 {
			return true
		}
	}

	return false
}
