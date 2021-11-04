package easy

func ClimbStairs(n int) int {
	arr := make([]int, n+1)

	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n]
}
