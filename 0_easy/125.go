package easy

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(s, "")
	processedString = strings.ToLower(processedString)

	for i := 0; i < len(processedString)/2; i++ {
		if processedString[i] != processedString[len(processedString)-1-i] {
			return false
		}
	}

	return true
}
