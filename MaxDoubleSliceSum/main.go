// https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/

package main

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	n := len(A)
	sumEndingAt := make([]int, n)
	sumStartingAt := make([]int, n)

	for i := 1; i < n-1; i++ {
		sumEndingAt[i] = maxInt(sumEndingAt[i-1]+A[i], 0)
		sumStartingAt[n-1-i] = maxInt(sumStartingAt[n-i]+A[n-1-i], 0)
	}

	maxDoubleSliceSum := 0
	for i := 1; i < len(A)-1; i++ {
		maxDoubleSliceSum = maxInt(maxDoubleSliceSum, sumEndingAt[i-1]+sumStartingAt[i+1])
	}
	return maxDoubleSliceSum

}

func main() {
	A := []int{3, 2, 6, -1, 4, 5, -1, 2}
	print(Solution(A))
}
