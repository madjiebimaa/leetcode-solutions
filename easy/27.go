package easy

func RemoveElement(nums []int, val int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = -1
			res++
		}
	}

	return res
}
