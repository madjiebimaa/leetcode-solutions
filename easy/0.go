package easy

import (
	"strings"
)

func BracketPair(s string) bool {
	arr := strings.Split(s, "")
	if len(arr) == 0 {
		return true
	}

	if arr[0] == "]" || arr[0] == ")" || arr[0] == "}" {
		return false
	}

	if len(arr)%2 != 0 {
		return false
	}

	m := make(map[string]int, len(s))
	m["("] = 0
	m["["] = 0
	m["{"] = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == "(" || arr[i] == "[" || arr[i] == "{" {
			m[arr[i]] += i
		}
	}

	for _, val := range m {
		if val%2 == 0 {
			return false
		}
	}
	// for i := 0; i < len(arr)/2; i++ {
	// 	if (arr[i] == "(" && arr[len(arr)-i-1] != ")") || (arr[i] == "[" && arr[len(arr)-i-1] != "]") || (arr[i] == "{" && arr[len(arr)-i-1] != "}") {
	// 		return false
	// 	}

	// }

	return true
}
