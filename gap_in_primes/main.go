// https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4

package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	w := 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}
	return true
}

func Gap(g, m, n int) []int {
	prime1 := 0
	prime2 := 0
	for i := m; i <= n; i++ {
		if IsPrime(i) {
			prime1, prime2 = prime2, i
			if prime1 != 0 && prime2-prime1 == g {
				return []int{prime1, prime2}
			}
		}

	}

	return nil
}

func main() {
	fmt.Println(Gap(2, 1, 10))
}
