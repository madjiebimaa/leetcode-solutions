package easy_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/madjiebimaa/go-leetcode/easy"
)

func TestTwoSum(t *testing.T) {
	res := easy.TwoSum([]int{1, 2, 3, 4}, 5)
	fmt.Println(res)
}

func TestIsPalindromeNumber(t *testing.T) {
	res := easy.IsPalindromeNumber(123)
	fmt.Println(res)
}

func TestRomanToInt(t *testing.T) {
	res := easy.RomanToInt("MCMXCIV")
	fmt.Println(res)
}

func TestSearchInsert(t *testing.T) {
	res := easy.SearchInsert([]int{1, 2, 4, 5, 6}, 0)
	fmt.Println(res)
}

func TestMaxSubArray(t *testing.T) {
	res := easy.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func TestLengthOfLastWord(t *testing.T) {
	res := strings.Fields("  foo bar  baz   ")
	fmt.Println(res)
}
