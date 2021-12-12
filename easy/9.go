package easy

import (
	"strconv"
	"strings"
)

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	xStr := strconv.Itoa(x)
	xSplit := strings.Split(xStr, "")
	for i := 0; i < len(xSplit); i++ {
		if xSplit[i] != xSplit[len(xSplit)-i-1] {
			return false
		}
	}

	return true
}
