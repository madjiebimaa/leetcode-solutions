package easy

func SearchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

		if i < len(nums)-1 {
			if nums[i] < target && nums[i+1] > target {
				return i + 1
			}
		}
	}

	return 0
}
