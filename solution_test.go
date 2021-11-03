package main

import (
	"testing"

	easy "github.com/madjiebimaa/go-leetcode/0_easy"
	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, easy.Generate(3), [][]int{{1}, {1, 1}, {1, 2, 1}}, "Equal")
	assert.Equal(t, easy.Generate(1), [][]int{{1}}, "Equal")

}
