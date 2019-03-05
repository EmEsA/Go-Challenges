package main

import "testing"

var tests = []struct {
	input    int
	expected string
}{
	{1, "white"},
	{5, "black"},
	{12, "black"},
	{42, "white"},
	{100, "black"},
	{2017, "white"},
}

func TestBlackOrWhiteKey(t *testing.T) {
	for _, test := range tests {
		if BlackOrWhiteKey(test.input) != test.expected {
			t.Errorf("key %v is expected to be %v", test.input, test.expected)
		}
	}

}
