// https://www.codewars.com/kata/56ed20a2c4e5d69155000301

package main

import (
	"fmt"
	"strings"
)

func Scale(s string, k, n int) string {
	if len(s) == 0 {
		return s
	}
	parts := strings.Split(s, "\n")

	endRes := []string{}
	for _, part := range parts {
		var tempRes strings.Builder
		for _, c := range part {
			for i := 0; i < k; i++ {
				tempRes.WriteString(string(c))
			}
		}
		for i := 0; i < n; i++ {
			endRes = append(endRes, tempRes.String())
		}
	}
	return strings.Join(endRes, "\n")
}

func main() {
	fmt.Println(Scale("", 1, 2))
}
