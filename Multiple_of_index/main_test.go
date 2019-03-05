package main

import "testing"

var tests = []struct {
	input  []int
	output []int
}{
	{[]int{22, -6, 32, 82, 9, 25}, []int{-6, 32, 25}},
	{[]int{68, -1, 1, -7, 10, 10}, []int{-1, 10}},
	{[]int{28, 38, -44, -99, -13, -54, 77, -51}, []int{38, -44, -99}},
}

func TestMultipleOfIndex(t *testing.T) {
	for _, test := range tests {
		res := multipleOfIndex(test.input)
		for i, v := range test.output {
			if v != res[i] {
				t.Error("Failed!")
			}
		}
	}
}
