package easy

func MaxSubArray(nums []int) int {
	sum := nums[0]
	if sum < 0 {
		sum = 0
	}

	max := nums[0]
	for _, v := range nums[1:] {
		sum = sum + v
		if sum > max {
			max = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return max
}
