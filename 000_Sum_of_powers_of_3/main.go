package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

const (
	n = 90
)

// Find all solutions of a^3 + b^3 = c^3 + d^3
// where a, b, c, d < n
func main() {

	var start time.Time
	var solutions int

	// O(n^4) solution
	solutions = 0
	start = time.Now()

	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			for c := 0; c < n; c++ {
				for d := 0; d < n; d++ {
					if math.Pow(float64(a), 3)+math.Pow(float64(b), 3) == math.Pow(float64(c), 3)+math.Pow(float64(d), 3) {
						solutions++
					}
				}
			}
		}
	}
	fmt.Println("Solutions O(n^4): ", solutions)
	log.Printf("Solutions O(n^4) took %s", time.Since(start))

	// O(n^2) solution
	solutions = 0
	start = time.Now()

	// First, precount the values and put it in a hash table (implemented as map in Go)
	precount := make(map[string]float64)
	for a := 0; a < n; a++ {
		for b := 0; b < n; b++ {
			key := strconv.Itoa(a) + "+" + strconv.Itoa(b)
			value := math.Pow(float64(a), 3) + math.Pow(float64(b), 3)
			precount[key] = value
		}
	}

	// Use only two loops to loop through precalculated values
	for _, value1 := range precount {
		for _, value2 := range precount {
			if value1 == value2 {
				solutions++
			}
		}
	}
	fmt.Println("Solutions O(n^2): ", solutions)
	log.Printf("Solutions O(n^2) took %s", time.Since(start))
}
