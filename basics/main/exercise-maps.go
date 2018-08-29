package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, v := range strings.Fields(s) {
		c, ok := res[v]
		if ok {
			res[v] = c + 1
		} else {
			res[v] = 1
		}
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
