// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_profit/

package main

import (
	"fmt"
	"math"
)

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	tempProfit := math.Max(0, float64(A[1])-float64(A[0]))
	maxProfit := tempProfit

	for i := 2; i < len(A); i++ {
		tempProfit = math.Max(0, tempProfit+float64(A[i]-A[i-1]))
		maxProfit = math.Max(maxProfit, tempProfit)
	}
	return int(maxProfit)
}

func main() {
	// A := []int{23171, 21011, 21123, 21366, 21013, 21367}
	A := []int{}
	fmt.Println(Solution(A))
}
