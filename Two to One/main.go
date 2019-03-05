package main

import (
	"fmt"
	"sort"
	"strings"
)

type LetterSet struct {
	letters map[string]struct{}
}

func (this *LetterSet) addFromString(s string) {
	for _, letter := range s {
		this.letters[string(letter)] = struct{}{}
	}
}

func (this *LetterSet) sorted() []string {
	letters := make([]string, len(this.letters))

	i := 0
	for letter := range this.letters {
		letters[i] = letter
		i++
	}
	sort.Strings(letters)
	return letters
}

func TwoToOne(s1 string, s2 string) string {
	set := LetterSet{make(map[string]struct{})}
	set.addFromString(s1)
	set.addFromString(s2)
	return strings.Join(set.sorted(), "")
}

// func main() {

// 	a := "xyaabbbccccdefww"
// 	b := "xxxxyyyyabklmopq"

// 	test := TwoToOne(a, b)
// 	fmt.Println(test == "abcdefklmopqwxy")

// }

func main() {
	fmt.Println([]byte("A"))
	fmt.Println([]byte("a"))

	a := "xyaAbbbccccdefww"
	b := "xxxxyEyyyDabCklmopq"

	test := TwoToOne(a, b)
	fmt.Println(test)

}
