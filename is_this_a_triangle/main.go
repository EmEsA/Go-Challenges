package main

import "fmt"

func IsTriangle(a, b, c int) bool {
	var max int
	for _, v := range []int{a, b, c} {
		if v > max {
			max = v
		}
	}
	if max >= a+b+c-max {
		return false
	}
	return true
}

func main() {
	fmt.Println(IsTriangle(1, 2, 3))
}
