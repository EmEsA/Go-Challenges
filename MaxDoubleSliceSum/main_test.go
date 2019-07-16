package main

import (
	"fmt"
	"testing"
)

type Test struct {
	input    []int
	expected int
}

var tests = []Test{
	{[]int{3, 2, 6, -1, 4, 5, -1, 2}, 17},
	{[]int{0, 2, 6, -1, 4, 5, 9}, 17},
	{[]int{3, 2, 1, -1, -1, -1, 4, 5, 8}, 10},
	{[]int{-1, -1, 4, 5, 8}, 9},
	{[]int{5, 17, 0, 3}, 17},
	{[]int{1, 1, -3, 4, 1}, 5},
	{[]int{1, -1, -3, 4, 1}, 4},
	{[]int{1, -1, -3}, 0},
	{[]int{0, 1, 2, -3, -1, 1, 2, 0}, 5},
}

func TestSolution(t *testing.T) {
	for _, test := range tests {
		res := Solution(test.input)
		if res != test.expected {
			t.Error(fmt.Sprintf("Failed! Expected %v, got %v", test.expected, res))
		}
	}
}
