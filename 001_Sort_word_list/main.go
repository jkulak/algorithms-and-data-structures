// Implement an algorithm to determine if a string has all unique runesPresent
// make it UTF ready
package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

type wordList []string

func (w wordList) Len() int {
	return len(w)
}

func (w wordList) Swap(a, b int) {
	w[a], w[b] = w[b], w[a]
}

func (w wordList) Less(a, b int) bool {
	// Converting to runes to make it UTF-8 safe
	return len([]rune(w[a])) < len([]rune(w[b]))
}

func (w wordList) Sort() wordList {
	sort.Sort(wordList(w))
	return w
}

func main() {

	var start time.Time
	start = time.Now()

	words := wordList{
		"strażnik",
		"teksasu",
		"Mexico",
		"Meksyk",
		"📙",
		"📙🔍",
		"jednolity",
		"najdłuższy",
		"powtórka",
		"🔮",
		"powtórka",
		"as",
		"łóżko",
		"Ewa",
		"Ola",
	}

	fmt.Println(words)
	fmt.Println(words.Sort())

	log.Printf("Time: %s", time.Since(start))
}
