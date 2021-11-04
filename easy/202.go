package easy

import (
	"fmt"
	"math"
	"strconv"
)

func IsHappy(n int) bool {
	if n <= 3 {
		return false
	}

	strNum := strconv.Itoa(n)
	var tempNum int
	for i := 0; i < len(strNum); i++ {
		num, _ := strconv.Atoi(strNum)
		tempNum += int(math.Pow(float64(num), 2))
		fmt.Println(tempNum)
		if i == len(strNum)-1 {
			if tempNum == 1 {
				return true
			} else if tempNum > 1 && tempNum <= 3 {
				return false
			}

			strNum = strconv.Itoa(tempNum)
			i = 0
		}
	}

	return false
}
