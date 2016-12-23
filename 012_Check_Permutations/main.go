// Implement an algorithm to determine if a string has all unique runesPresent
// make it UTF ready
package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func isPermutation(a, b string) bool {

	c := []rune(a)
	d := []rune(b)

	if len(a) != len(b) {
		return false
	}

	for _, v := range c {
		for i := 0; i < len(d); i++ {
			if d[i] == v {
				// create new slice without the element at i
				d = append(d[:i], d[i+1:]...)
				break
			}
		}
	}

	if len(d) == 0 {
		return true
	}

	return false
}

// Define a type that is an array of runes
type runeList []rune

// Add Less function to the type
func (s runeList) Less(i, j int) bool {
	return s[i] < s[j]
}

// Add Swap method to the type
func (s runeList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Add Len method to the type
func (s runeList) Len() int {
	return len(s)
}

func sortRuneList(s string) string {
	r := []rune(s)
	sort.Sort(runeList(r))
	return string(r)
}

func isPermutationSort(a, b string) bool {

	if len(a) != len(b) {
		return false
	}

	c := sortRuneList(a)
	d := sortRuneList(b)

	if c == d {
		return true
	}

	return false
}

func main() {

	var start time.Time
	start = time.Now()

	type testData struct {
		a   string
		b   string
		res bool
	}

	// Three alternative ways to declate and initiate the variables

	// var tests []testData
	// tests = make([]testData, 10)
	// tests[0] = testData{"a", "a", true}
	// tests[1] = testData{"as", "sa", true}

	// var tests [10]testData
	// tests[0] = testData{"a", "a", true}
	// tests[1] = testData{"as", "sa", true}

	// tests := []testData{
	// 	testData{"a", "a", true},
	// 	testData{"as", "sa", true},
	// }

	var tests = []testData{
		testData{"a", "a", true},
		testData{"as", "sa", true},
		testData{"asa", "saa", true},
		testData{"asaa", "ssaa", false},
		testData{"aSaa", "saaa", false},
		testData{"asaa ", "ssaaa", false},
		testData{"ass", "saa", false},
		testData{"asss", "sssa", true},
		testData{"🚙🚜🚜🚗", "🚙🚜🚜", false},
		testData{"🚙🚜🚜🚗", "sa", false},
		testData{"🚙🚜🚜🚗", "🚙🚜🚜🚗🚙🚜🚜🚗", false},
		testData{"🚙🚜🚜🚗🚙🚜🚜🚗", "🚙🚜🚜🚗", false},
		testData{"🚙🚜🚜🚗🚙🚜🚜🚗", "🚜🚜🚙🚙🚗🚜🚜🚗", true},
		testData{"🚙🚙łł", "łł🚙🚙", true},
		testData{"a🚙🚙łł", "łł🚙🚙", false},
		testData{"a🚙🚙łł", "łła🚙🚙", true},
		testData{"śća🚙🚙łł", "łła🚙🚙", false},
		testData{"śća🚙🚙łł", "łła🚙ść🚙", true},
		testData{"🎸🎸", "🎸", false},
		testData{"🎸🎸", "🎸🎸🎸🎸", false},
		testData{"🎸🎸🎸🎸", "🎸🎸🎸🎸", true},
		testData{"🎸🎸🎸💉", "🎸🎸🎸🎸", false},
	}

	fmt.Printf("%v\n", tests)

	for key, val := range tests {
		res := isPermutationSort(val.a, val.b)

		if res != val.res {
			fmt.Printf("🚫  ")
		} else {
			fmt.Printf("✅  ")
		}
		fmt.Printf("Test %d - %s, %s (expected: %v, returned: %v)\n", key+1, val.a, val.b, val.res, res)
	}

	log.Printf("Time: %s", time.Since(start))
}
