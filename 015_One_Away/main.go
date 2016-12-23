// Write a method to replace all spaces in a string with "%20"
// make it UTF ready
package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

type testData struct {
	a   string
	b   string
	res bool
}

func logTest(key int, val testData, res bool, verbose bool) {
	if res != val.res {
		fmt.Printf("😢 ")
	} else {
		fmt.Printf("✅ ")
	}

	if verbose {
		fmt.Printf(" test %2d exp: \"%v\" got: \"%v\"\t%v\n", key, val.res, res, val)
	}
}

func editedOne(a, b string) bool {
	lena := len(a)
	dif := false

	for i := 0; i < lena; i++ {
		if a[i] != b[i] {
			if dif {
				return false
			}
			dif = true
		}
	}
	return true
}

func addedOne(a, b string) bool {

	lena, lenb := len(a), len(b)
	index1, index2 := 0, 0

	for index1 < lena && index2 < lenb {
		if a[index1] != b[index2] {
			if index1 != index2 {
				return false
			}
			index1++
		} else {
			index1++
			index2++
		}
	}
	return true
}

func checkIfOneAway(a, b string) bool {

	lena, lenb := len(a), len(b)

	// Length difference can't be bigger than 1
	if math.Abs(float64(lena-lenb)) > 1 {
		return false
	}

	// Check if strings are not the same
	if a == b {
		return false
	}

	// Remove one character
	if lena > lenb {
		return addedOne(a, b)
	}

	// Add one character
	if lena < lenb {
		return addedOne(b, a)
	}

	// Replace one character
	return editedOne(a, b)
}

func main() {

	paramVerbose := flag.Bool("v", false, "be verbose")
	paramTestNumber := flag.Int("t", -1, "run selected test only")
	flag.Parse()

	var verbose bool = *paramVerbose
	var testNumber int = *paramTestNumber

	var start time.Time = time.Now()

	var tests = []testData{

		testData{"tac", "otac", true},
		testData{"tac", "toac", true},
		testData{"tac", "taoc", true},
		testData{"tac", "taco", true},
		testData{"tac", "ac", true},
		testData{"tac", "tc", true},
		testData{"tac", "ta", true},
		testData{"tac", "xac", true},
		testData{"tac", "txc", true},
		testData{"tac", "tax", true},

		testData{"aaa", "aaaa", true},
		testData{"aaa", "aaac", true},
		testData{"aaa", "caaa", true},
		testData{"aaa", "acaa", true},
		testData{"aaa", "aaca", true},
		testData{"aaa", "aa", true},
		testData{"aaa", "gaa", true},
		testData{"aaa", "aga", true},
		testData{"aaa", "aag", true},

		// no changes
		testData{"tac", "tac", false},
		testData{"tac", "zztac", false},
		testData{"tac", "tzzac", false},
		testData{"tac", "tazzc", false},
		testData{"tac", "taczz", false},

		testData{"tac", "t", false},
		testData{"tac", "c", false},
		testData{"tac", "a", false},

		testData{"tac", "nnac", false},
		testData{"tac", "tnnc", false},
		testData{"tac", "tann", false},

		testData{"tac", "atc", false},
		testData{"tac", "tca", false},
		testData{"tac", "cat", false},

		testData{"tac", "av", false},
		testData{"tac", "va", false},
		testData{"tac", "vt", false},
		testData{"tac", "tv", false},
		testData{"tac", "cv", false},
		testData{"tac", "vc", false},

		testData{"aaa", "aaaaa", false},
		testData{"aaa", "aaaab", false},
		testData{"aaa", "aaaba", false},

		testData{"aaa", "acc", false},
		testData{"aaa", "cca", false},
		testData{"aaa", "cac", false},

		testData{"aaa", "ac", false},
		testData{"aaa", "ca", false},
		testData{"aaa", "a", false},
		testData{"aaa", "a", false},

		testData{"tac", "💙", false},
		testData{"tac", "💙💙💙💙", false},
	}

	if testNumber > 0 {
		val := tests[testNumber]
		res := checkIfOneAway(val.a, val.b)
		key := testNumber
		logTest(key, val, res, verbose)
	} else {
		for key, val := range tests {
			res := checkIfOneAway(val.a, val.b)
			logTest(key, val, res, verbose)
		}
	}

	fmt.Printf("\nTime: %s\n\n", time.Since(start))
}
