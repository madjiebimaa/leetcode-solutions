package easy

import (
	"math"
)

func MaxProfit(prices []int) int {
	min := int(math.Pow10(4))
	max := 0

	for i := 0; i < len(prices); i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		max = int(math.Max(float64(max), float64(prices[i]-min)))
	}

	return max
}
