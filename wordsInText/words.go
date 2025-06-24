package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Привет мир привет Go мир Go Go"
	fmt.Println(text)

	wordFrequency := countWords(text)

	fmt.Println("Частота слов:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}
