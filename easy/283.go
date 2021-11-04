package easy

func MoveZeroes(nums []int) []int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			cnt++
		}
	}

	for i := cnt; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}
