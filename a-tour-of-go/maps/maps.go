package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	arr := strings.Fields(s)
	for _, e := range arr {
		m[e] = m[e] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
