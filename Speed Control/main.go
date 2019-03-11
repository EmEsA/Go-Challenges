// https://www.codewars.com/kata/56484848ba95170a8000004d

package main

import "fmt"

func Gps(s int, x []float64) int {
	vMax := 0
	for i := 0; i < len(x)-1; i++ {
		if v := int((3600 * (x[i+1] - x[i])) / float64(s)); v > vMax {
			vMax = v
		}
	}
	return vMax
}

func main() {
	x := []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}
	// x := []float64{}
	s := 15

	fmt.Println(Gps(s, x))
}
