// Write a method to replace all spaces in a string with "%20"
// make it UTF ready
package main

import (
	"fmt"
	"time"
)

const (
	debug = false
)

// O(n) version
func checkIfPalindromePermutationOptimized(s string) bool {
	a := ""
	length := len(s)

	// Remove spaces O(n)
	for i := 0; i < length; i++ {
		if string(s[i]) != " " {
			a = a + string(s[i])
		}
	}

	var chars map[byte]int
	chars = make(map[byte]int)

	length = len(a)
	for i := 0; i < length; i++ {
		chars[a[i]]++
	}

	foundOdd := false
	for _, v := range chars {
		if v%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}

// O(n^2) version
func checkIfPalindromePermutation(s string) bool {
	a := ""
	length := len(s)

	// Remove spaces O(n)
	for i := 0; i < length; i++ {
		if string(s[i]) != " " {
			a = a + string(s[i])
		}
	}

	// Remove pairs of characters O(n^2)
	for m := 0; m < len(a); m++ {
		for n := m + 1; n < len(a); n++ {
			if a[m] == a[n] {
				a = a[:n] + a[n+1:]
				a = a[:m] + a[m+1:]
				m--
				break
			}
		}
	}

	if len(a) < 2 {
		return true
	}

	return false
}

func main() {

	var start time.Time
	start = time.Now()

	type testData struct {
		a   string
		res bool
	}

	var tests = []testData{
		// maka akam
		testData{"tac", false},
		testData{"mmaakk aa", true},
		// taco cat
		testData{"caot atc", true},
		testData{"coat atc", true},
		testData{"cott aac", true},
		testData{"coto aac", true},
		testData{"coto acc", false},
		testData{"coto acbc", false},
		testData{"coto acbcc", false},
		testData{"coto acbxc", false},
		testData{"coto zacbxc", false},
	}

	if debug {
		fmt.Printf("%v\n", tests)
	}

	for key, val := range tests {
		// res := checkIfPalindromePermutation(val.a)
		res := checkIfPalindromePermutationOptimized(val.a)

		if res != val.res {
			fmt.Printf("🚫  ")
		} else {
			fmt.Printf("✅  ")
		}

		if debug {
			fmt.Printf("Test %d %v\n\texpected: \"%t\"\n\treturned: \"%t\"\n", key+1, val, val.res, res)
		}
	}
	fmt.Println()
	fmt.Printf("Time: %s", time.Since(start))
}
