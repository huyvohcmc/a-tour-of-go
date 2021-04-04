package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	res := make(map[string]int)
	for _, word := range words {
		if _, ok := res[word]; ok {
			res[word] += 1
		} else {
			res[word] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
