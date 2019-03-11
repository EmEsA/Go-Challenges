package main

import (
	"testing"
)

type Test struct {
	input  int
	output int64
}

var SuMintests = []Test{
	{5, 55},
	{100, 338350},
}

var SuMaxtests = []Test{
	{5, 95},
	{100, 671650},
}

func runTest(t *testing.T, f func(n int) int64, funcName string, testCases []Test) {
	for _, test := range testCases {
		res := f(test.input)
		if res != test.output {
			t.Errorf("%s(%v) = %v; want %v", funcName, test.input, res, test.output)
		}
	}
}

func TestSuMin(t *testing.T) { runTest(t, SuMin, "SuMin", SuMintests) }
func TestSuMax(t *testing.T) { runTest(t, SuMax, "SuMax", SuMaxtests) }
