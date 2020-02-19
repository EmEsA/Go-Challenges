// https://www.codewars.com/kata/56dbe7f113c2f63570000b86

package main

import (
	"fmt"
	"strings"
)

func Rot(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(s))
	for i := 0; i < len(runes); i++ {
		reversed[len(runes)-1-i] = runes[i]
	}
	return string(reversed)
}

func SelfieAndRot(s string) string {
	dots := strings.Repeat(".", strings.Index(s, "\n"))
	sWithDots := strings.ReplaceAll(s, "\n", dots+"\n")
	return sWithDots + dots + "\n" + dots + Rot(sWithDots)
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func main() {
	inp1 := "abcd\nefgh\nijkl\nmnop"
	exp1 := "ponm\nlkji\nhgfe\ndcba"
	res1 := Oper(Rot, inp1)
	fmt.Println(res1 == exp1)

	inp2 := "xZBV\njsbS\nJcpN\nfVnP"
	exp2 := "xZBV....\njsbS....\nJcpN....\nfVnP....\n....PnVf\n....NpcJ\n....Sbsj\n....VBZx"
	res2 := Oper(SelfieAndRot, inp2)
	fmt.Println(res2 == exp2)
}
