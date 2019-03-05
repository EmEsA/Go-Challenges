package main

import "fmt"

func BlackOrWhiteKey(keyPressCount int) string {
	n := (keyPressCount-1)%88 + 1 // excercise loops after 88 keys
	n = (n-1)%12 + 1 // key pattern loops after 12 keys
	if n == 2 || n == 5 || n == 7 || n == 10 || n == 12 {
		return "black"
	} else {
		return "white"
	}
}

func main() {
	fmt.Println(BlackOrWhiteKey(1))

}
