// https://www.codewars.com/kata/5868b2de442e3fb2bb000119

package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Element struct {
	Index  int
	Number int
	Weight int
}

func Closest(strng string) string {
	if strng == "" {
		return "[(), ()]"
	}
	var data []Element
	for i, str := range strings.Split(strings.Trim(strng, " "), " ") {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Element{
			Index:  i,
			Number: num,
			Weight: weight(num),
		})
	}
	sort.Slice(data, func(i, j int) bool {
		if data[i].Weight != data[j].Weight {
			return data[i].Weight < data[j].Weight
		}
		return data[i].Index < data[j].Index
	})

	var minDiffInd int
	minDiff := -1
	for i := 0; i < len(data)-1; i++ {
		diff := data[i+1].Weight - data[i].Weight
		if diff < minDiff || minDiff < 0 {
			minDiff = diff
			minDiffInd = i
		}
	}
	return fmt.Sprintf("[(%v, %v, %v), (%v, %v, %v)]",
		data[minDiffInd].Weight, data[minDiffInd].Index, data[minDiffInd].Number,
		data[minDiffInd+1].Weight, data[minDiffInd+1].Index, data[minDiffInd+1].Number)

}

func weight(n int) int {
	var w int
	for n > 0 {
		w += n % 10
		n = n / 10
	}
	return w
}

func main() {
	inp := "239382 162 254765 182 485944 134 468751 62 49780 108 54"
	exp := "[(8, 5, 134), (8, 7, 62)]"
	fmt.Println(exp == Closest(inp))
}
