package easy

func MissingNumber(nums []int) int {
	countAll, countArr := len(nums), 0
	for i := 0; i < len(nums); i++ {
		countAll += i
		countArr += nums[i]
	}

	return countAll - countArr
}
