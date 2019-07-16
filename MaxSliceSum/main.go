package main

import "math"

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	currSlice := float64(A[0])
	maxSlice := currSlice

	for i := 1; i < len(A); i++ {
		a := float64(A[i])
		currSlice = math.Max(currSlice+a, a)
		maxSlice = math.Max(currSlice, maxSlice)
	}
	return int(maxSlice)
}

func main() {
	// A := []int{3, 2, -6, 4, 0}
	A := []int{-1, 1, 0}
	println(Solution(A))
}
