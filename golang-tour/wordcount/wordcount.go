package main

import (
	"strings"
	"golang.org/x/tour/wc"
)


// WordCount function prints number of occurences for each word
func WordCount(s string) map[string]int {
	wordList := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	testString := "Hi I'm Hide Hi"
	WordCount(testString)
	// wc.Test(WordCount)
}
