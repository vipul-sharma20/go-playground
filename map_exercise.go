package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	a := strings.Fields(s)
	for _, e := range a {
		_, ok := m[e]
		if ok {
			m[e]++
		} else {
			m[e] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

