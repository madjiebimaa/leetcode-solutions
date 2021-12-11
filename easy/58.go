package easy

import "strings"

func LengthOfLastWord(s string) int {
	arrS := strings.Fields(s)
	return len(arrS[len(arrS)-1])
}
