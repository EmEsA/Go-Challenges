package main

//SuMin returns sum of mins
func SuMin(m int) int64 {
	var sum int
	for i, j := m, 1; i >= 1; i, j = i-1, j+2 {
		sum += i * j
	}
	return int64(sum)
}

//SuMax return sum of maxs
func SuMax(m int) int64 {
	var sum int
	for i, j := 1, 1; i <= m; i, j = i+1, j+2 {
		sum += i * j
	}
	return int64(sum)
}

//SumSum returns sum of both
func SumSum(m int) int64 {
	return SuMin(m) + SuMax(m)
}

func main() {

}
