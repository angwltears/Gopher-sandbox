package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apple banana apple orange banana apple"
	words := strings.Fields(str)

	counts := make(map[string]int)

	for _, word := range words {

		counts[word]++
	}

	fmt.Println(counts)
}
