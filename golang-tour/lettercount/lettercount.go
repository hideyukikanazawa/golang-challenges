package main

import (
	"fmt"
	"regexp"
	"strings"
)

func letterCount(s string) map[rune]int {
	letters := make(map[rune]int)
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-zA-Z]")
	alphaS := reg.ReplaceAllString(s, "")
	for _, v := range alphaS {
		if _, ok := letters[v]; ok {
			letters[v]++
		} else {
			letters[v] = 1
		}
	}
	return letters
}

func main() {
	s := "Barbotoss"
	for letter, count := range letterCount(s) {
		fmt.Printf("%v: %v times\n", string(letter), count)
	}
}
