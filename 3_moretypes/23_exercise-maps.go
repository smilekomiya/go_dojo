package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	stringArray := strings.Fields(s)
	for i := 0; i < len(stringArray); i++ {
		count, ok := m[stringArray[i]]
		if ok {
			m[stringArray[i]] = count + 1
		} else {
			m[stringArray[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
