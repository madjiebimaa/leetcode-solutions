package easy

func MajorityElement(nums []int) int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			m[nums[i]] = 1
		} else {
			m[nums[i]] += 1
		}
	}

	res := 0
	vMax := 0
	for k, v := range m {
		if v > vMax {
			res = k
			vMax = v
		}
	}

	return res
}
