package main

import "testing"

type Test struct {
	input    []int
	expected bool
}

var tests = []Test{
	{[]int{5, 1, 2}, false},
	{[]int{1, 2, 5}, false},
	{[]int{2, 5, 1}, false},
	{[]int{4, 2, 3}, true},
	{[]int{5, 1, 5}, true},
	{[]int{2, 2, 2}, true},
	{[]int{-1, 2, 3}, false},
}

func TestIsTriangle(t *testing.T) {
	for _, test := range tests {
		res := IsTriangle(test.input[0], test.input[1], test.input[2])
		if res != test.expected {
			t.Fatalf("IsTriangle(%v, %v, %v) = %v, expected: %v", test.input[0], test.input[1], test.input[2], res, test.expected)
		}
	}
}
