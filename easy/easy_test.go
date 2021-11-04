package easy_test

import (
	"testing"

	. "github.com/madjiebimaa/go-leetcode/easy"
	"github.com/stretchr/testify/assert"
)

func TestSolution344(t *testing.T) {
	testTable := []struct {
		Input    []byte
		Expected []byte
	}{
		{Input: []byte{'h', 'e', 'l', 'l', 'o'}, Expected: []byte{'o', 'l', 'l', 'e', 'h'}},
		{Input: []byte{'a', 'd', 'j', 'i', 'e'}, Expected: []byte{'e', 'i', 'j', 'd', 'a'}},
	}

	for _, test := range testTable {
		assert.Equal(t, ReverseString(test.Input), test.Expected)
	}
}

func TestSolution326(t *testing.T) {
	testTable := []struct {
		Input    int
		Expected bool
	}{
		{Input: -2, Expected: false},
		{Input: 0, Expected: false},
		{Input: 1, Expected: true},
		{Input: 3, Expected: true},
		{Input: 6, Expected: false},
	}

	for _, test := range testTable {
		assert.Equal(t, IsPowerOfThree(test.Input), test.Expected)
	}
}

func TestSolution412(t *testing.T) {
	testTable := []struct {
		Input    int
		Expected []string
	}{
		{Input: 3, Expected: []string{"1", "2", "Fizz"}},
		{Input: 5, Expected: []string{"1", "2", "Fizz", "4", "Buzz"}},
	}

	for _, test := range testTable {
		assert.Equal(t, FizzBuzz(test.Input), test.Expected)
	}

}

func TestSolution387(t *testing.T) {
	testTable := []struct {
		Input    string
		Expected int
	}{
		{Input: "leetcode", Expected: 0},
		{Input: "loveleetcode", Expected: 2},
		{Input: "aabb", Expected: -1},
	}

	for _, test := range testTable {
		assert.Equal(t, FirstUniqChar(test.Input), test.Expected)
	}

}

func TestSolution283(t *testing.T) {
	testTable := []struct {
		Input    []int
		Expected []int
	}{
		{Input: []int{1, 2, 3, 0, 2, 3}, Expected: []int{1, 2, 3, 2, 3, 0}},
		{Input: []int{0}, Expected: []int{0}},
		{Input: []int{0, 1, 0, 3, 12}, Expected: []int{1, 3, 12, 0, 0}},
	}

	for _, test := range testTable {
		assert.Equal(t, MoveZeroes(test.Input), test.Expected)
	}
}

func TestSolution268(t *testing.T) {
	testTable := []struct {
		Input    []int
		Expected int
	}{
		{Input: []int{3, 0, 1}, Expected: 2},
		{Input: []int{0, 1, 4, 5, 6, 3}, Expected: 2},
		{Input: []int{1, 2, 3, 4}, Expected: 0},
	}

	for _, test := range testTable {
		assert.Equal(t, MissingNumber(test.Input), test.Expected)
	}
}

func TestSolution202(t *testing.T) {
	t.Skip("Not solved")
	testTable := []struct {
		Input    int
		Expected bool
	}{
		{Input: 19, Expected: true},
		{Input: 2, Expected: false},
		{Input: 3, Expected: false},
	}

	for _, test := range testTable {
		assert.Equal(t, IsHappy(test.Input), test.Expected)
	}
}

// ------- Wrong place for positioning actual and expected
func TestSolution20(t *testing.T) {
	testTable := []struct {
		Input    string
		Expected bool
	}{
		{Input: "", Expected: true},
		{Input: "][", Expected: false},
		{Input: "()", Expected: true},
		{Input: "([)", Expected: false},
		{Input: "{[]}", Expected: true},
	}

	for _, test := range testTable {
		assert.Equal(t, test.Expected, IsValid(test.Input))
	}
}

func TestSolution26(t *testing.T) {
	testTable := []struct {
		Input    []int
		Expected int
	}{
		{Input: []int{1, 1, 2}, Expected: 2},
		{Input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, Expected: 5},
	}

	for _, test := range testTable {
		assert.Equal(t, test.Expected, RemoveDuplicates(test.Input))
	}
}
