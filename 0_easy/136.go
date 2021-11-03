package easy

func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			m[nums[i]] = 1
		} else {
			m[nums[i]]++
		}
	}

	for key, val := range m {
		if val == 1 {
			return key
		}
	}

	return -1
}
