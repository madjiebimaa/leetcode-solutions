package easy_test

import (
	"fmt"
	"testing"

	"github.com/madjiebimaa/go-leetcode/easy"
)

func TestMajorityElement(t *testing.T) {
	res := easy.MajorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(res)
}

func TestIntersect(t *testing.T) {
	res := easy.Intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(res)
}
