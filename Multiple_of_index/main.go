package main

func multipleOfIndex(ints []int) []int {
	res := []int{}
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			res = append(res, ints[i])
		}
	}
	return res
}

func main() {

}
