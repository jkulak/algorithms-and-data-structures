package main

import (
	"math/rand"
	"time"
)

// InitArray inits an array of random integers
func InitArray(size, max int) []int {

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	var a = make([]int, size)

	// Create an array with random values
	for i := 0; i < size; i++ {
		a[i] = r.Intn(max)
	}

	return a
}
